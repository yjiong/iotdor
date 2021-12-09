package datasrc

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	mqtt "github.com/eclipse/paho.mqtt.golang"

	//"github.com/eclipse/paho.mqtt.golang/packets"
	log "github.com/sirupsen/logrus"
)

// DataDownPayload ...
type DataDownPayload struct {
	Pj *simplejson.Json
}

// MQTTDSrcer ...
type MQTTDSrcer struct {
	conn         mqtt.Client
	dataDownChan chan DataDownPayload
	wg           sync.WaitGroup
	Iotdname     string
	onlinemsg    string
	pqos         byte
	sqos         byte
	retain       bool
	SubTopic     string
}

// NewMQTTDSrcer creates a new MQTTDSrcer.
func NewMQTTDSrcer(conm map[string]string) DSrcer {
	h := MQTTDSrcer{
		dataDownChan: make(chan DataDownPayload),
	}

	opts := mqtt.NewClientOptions()
	server := conm["url"]
	h.Iotdname = conm["iotdname"]
	log.Debugln(conm)
	opts.SetUsername(conm["username"])
	opts.SetPassword(conm["password"])
	opts.SetOnConnectHandler(h.onConnected)
	opts.SetConnectionLostHandler(h.onConnectionLost)
	opts.AutoReconnect = true
	kplv, _ := strconv.Atoi(conm["keepalive"])
	pqos, _ := strconv.Atoi(conm["pqos"])
	h.pqos = byte(pqos)
	sqos, _ := strconv.Atoi(conm["sqos"])
	h.sqos = byte(sqos)
	h.retain, _ = strconv.ParseBool(conm["retain"])
	opts.SetKeepAlive(time.Duration(kplv) * time.Second)
	opts.SetClientID(h.Iotdname)
	optserver := "tcp://" + server
	if conm["cafile"] != "" {
		tlsconfig, err := newTLSConfig(conm)
		if err != nil {
			log.Errorf("Error with the mqtt CA certificate: %s", err)
		} else {
			opts.SetTLSConfig(tlsconfig)
			optserver = "ssl://" + server
		}
	}
	opts.AddBroker(optserver)

	log.WithField("server", server).Info("handler/mqtt: connecting to mqtt broker")
	h.conn = mqtt.NewClient(opts)
	var counts int
	for {
		if token := h.conn.Connect(); token.Wait() && token.Error() != nil {
			if counts < 10 {
				log.Errorf("handler/mqtt: connecting to broker error, will retry in 2s: %s", token.Error())
				time.Sleep(2 * time.Second)
				counts++
			} else {
				log.Errorf("handler/mqtt: connecting to broker error, will retry in 30s: %s", token.Error())
				time.Sleep(60 * time.Second)
			}
		} else {
			log.Info("handler/mqtt: conneting successfull")
			break
		}
	}
	return &h
}

func newTLSConfig(cm map[string]string) (*tls.Config, error) {
	// Import trusted certificates from CAfile.pem.
	cafile := cm["cafile"]
	cert, err := ioutil.ReadFile(cafile)
	//log.Infof("print this ca file :\n %s", cert)
	if err != nil {
		log.Errorf("backend: couldn't load cafile: %s", err)
		return nil, err
	}

	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(cert)

	// Create tls.Config with desired tls properties
	if cm["certfile"] != "" && cm["keyfile"] != "" {
		certpair, err := tls.LoadX509KeyPair(cm["certfile"], cm["keyfile"])
		if err != nil {
			log.Errorf("get cert error :%s", err)
			return nil, err
			//log.Fatalf("get cert error :%s", err)
		}
		return &tls.Config{
			RootCAs:            certpool,
			Certificates:       []tls.Certificate{certpair},
			InsecureSkipVerify: true,
		}, nil
	}
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs:            certpool,
		InsecureSkipVerify: true,
	}, nil
}

// Close stops the handler.
func (h *MQTTDSrcer) Close() error {
	log.Info("handler/mqtt: closing handler")
	if token := h.conn.Unsubscribe(h.Iotdname + "/+"); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: unsubscribe from %s error: %s", h.Iotdname, token.Error())
	}
	log.Info("handler/mqtt: handling last items in queue")
	h.wg.Wait()
	close(h.dataDownChan)
	return nil
}

// SendData sends a DataUpPayload.
func (h *MQTTDSrcer) SendData(tp string, payload interface{}) error {
	b, err := json.Marshal(payload)
	topic := tp
	if err != nil {
		return fmt.Errorf("handler/mqtt: data-up payload marshal error: %s", err)
	}
	if token := h.conn.Publish(topic, h.pqos, h.retain, b); token.Wait() && token.Error() != nil {
		return fmt.Errorf("handler/mqtt: publish data-up error: %s", err)
	}
	mymsg, err := simplejson.NewJson(b)
	logsb, _ := mymsg.EncodePretty()
	log.WithFields(log.Fields{"topic": topic, "messageType": "handler/mqtt: publishing data-send"}).Debugf(string(logsb))
	return nil
}

// DataDownChan returns the channel containing the received DataDownPayload.
func (h *MQTTDSrcer) DataDownChan() chan DataDownPayload {
	return h.dataDownChan
}

func (h *MQTTDSrcer) rxmsgDSrcer(c mqtt.Client, msg mqtt.Message) {
	h.wg.Add(1)
	defer h.wg.Done()
	mymsgjson, err := simplejson.NewJson(msg.Payload())
	if err != nil {
		log.WithFields(log.Fields{
			"msg_payload": string(msg.Payload()),
		}).Errorf("message is not json format: %s", err)
		return
	}
	logsb, _ := mymsgjson.EncodePretty()
	log.WithFields(log.Fields{"topic": msg.Topic(), "messageType": "received cmd", "Qos": msg.Qos()}).Info(string(logsb))
	h.dataDownChan <- DataDownPayload{Pj: mymsgjson}
}

func (h *MQTTDSrcer) onConnected(c mqtt.Client) {
	log.Infof("handler/mqtt: connected to mqtt broker subscribe qos=%d, Publish qos=%d", h.sqos, h.pqos)
	tps := []string{h.Iotdname + "/+"}
	if h.SubTopic != "" {
		tps = append(tps, h.SubTopic)
	}
	log.Debugf("sub_topic=%s", tps)
	for _, tp := range tps {
		for {
			log.WithField("topic", tp).Info("handler/mqtt: subscribling to iot topic")
			if token := h.conn.Subscribe(tp, h.sqos, h.rxmsgDSrcer); token.Wait() && token.Error() != nil {
				log.WithField("topic", tp).Errorf("handler/mqtt: subscribe error: %s", token.Error())
				time.Sleep(time.Second)
				continue
			}
			break
		}
	}
}

func (h *MQTTDSrcer) onConnectionLost(c mqtt.Client, reason error) {
	log.Errorf("handler/mqtt: mqtt connection error: %s", reason)
	//common.Mqttconnected = false
}

//IsConnected ..
func (h *MQTTDSrcer) IsConnected() bool {
	return h.conn.IsConnected()
}

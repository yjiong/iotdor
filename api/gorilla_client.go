package api

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//TranData ...
type TranData struct {
	UUID   string      `json:"uuid"`
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
	RawMsg []byte      `json:"raw_msg"`
}

//WSClient ...
type WSClient struct {
	conn    *websocket.Conn
	sn      string
	url     string
	errors  chan error
	stop    chan struct{}
	inMsg   chan TranData
	outMsg  chan string
	cliHTTP net.Conn
}

// InitWSClint ...
//url:"ws://api/ws/remoteConfig"
func InitWSClint(sn, url string) *WSClient {
	ws := &WSClient{
		conn:   nil,
		sn:     sn,
		url:    url,
		errors: make(chan error),
		stop:   make(chan struct{}),
		inMsg:  make(chan TranData),
		outMsg: make(chan string),
	}
	ws.connect()
	go ws.errorWatch()
	go ws.listenChan()
	//go ws.receiveMessage()
	return ws
}

func (w *WSClient) connect() {
	for {
		conn, _, err := websocket.DefaultDialer.Dial(w.url, http.Header{"sn": []string{w.sn}})
		if err != nil {
			<-time.After(time.Second * 2)
			continue
		}
		log.Debugf("websocket %s on connected", conn.LocalAddr().String())
		w.conn = conn
		for {
			if w.cliHTTP != nil {
				w.cliHTTP.Close()
			}
			if clihttp, err := net.Dial("tcp", "127.0.0.1:8888"); err == nil {
				log.Debugf("%s on connected", clihttp.LocalAddr().String())
				wconn := w.conn.UnderlyingConn()
				w.cliHTTP = clihttp
				go w.ioCopyPrint(w.cliHTTP, wconn)
				go w.ioCopyPrint(wconn, w.cliHTTP)
				return
			}
			<-time.After(time.Second * 2)
			log.Error(err)
			continue
		}
	}
}

func (w *WSClient) ioCopyPrint(dst, src net.Conn, isP ...bool) {
	for {
		select {
		case <-w.stop:
			return
		default:
			bf := make([]byte, (32 << 10))
			if l, er := src.Read(bf); er == nil {
				if len(isP) > 0 && isP[0] {
					log.Debugf("\nsrc:%s---->dst:%s \ncontent: %s\nhex content: % X",
						src.LocalAddr().String(),
						dst.LocalAddr().String(),
						bf[:l], bf[:l],
					)
				}
				dst.Write(bf[:l])
			} else {
				err := fmt.Errorf("ioCopyPrint func out \nsrc:%s---->dst:%s  %s",
					src.LocalAddr().String(),
					dst.LocalAddr().String(),
					er)
				log.Debugln(err)
				w.singalError(err)
				return
			}
		}
	}
}

func (w *WSClient) reConnection() {
	w.connect()
	w.stop = make(chan struct{})
	go w.errorWatch()
	go w.listenChan()
	//go w.receiveMessage()
}

func (w *WSClient) singalError(err error) {
	select {
	case w.errors <- err:
	default:
	}
}

func (w *WSClient) errorWatch() {
	select {
	case <-w.stop:
		return
	case err := <-w.errors:
		log.Warning(errors.Wrap(err, "in wsclient errorWatch"))
		close(w.stop)
		return
	}
}

func (w *WSClient) receiveMessage() {
	for {
		select {
		case <-w.stop:
			return
		default:
			_, p, err := w.conn.ReadMessage()
			if err != nil {
				w.singalError(err)
				return
			}
			if data, err := MsgDeCompress(p); err == nil {
				w.inMsg <- data
			} else {
				log.Errorln(err)
			}
		}
	}
}

// SendMessage ...
func (w *WSClient) SendMessage(msg TranData) {
	if data, err := MsgCompress(msg); err == nil {
		w.outMsg <- data
	}
}

func (w *WSClient) listenChan() {
	for {
		select {
		case <-w.stop:
			w.reConnection()
			return
			//case inMsg := <-w.inMsg:
			//go w.msgProcess(inMsg)
			//case outMsg := <-w.outMsg:
			//if err := w.conn.WriteMessage(websocket.BinaryMessage, []byte(outMsg)); err != nil {
			//w.singalError(err)
			//}
		}
	}
}

// MsgDeCompress decompress bytes
func MsgDeCompress(dataBytes []byte) (TranData, error) {
	var obj TranData

	tmp, err := base64.URLEncoding.DecodeString(string(dataBytes))
	if err != nil {
		return obj, err
	}
	zlibr, _ := zlib.NewReader(bytes.NewBuffer(tmp))
	dec := json.NewDecoder(zlibr)
	dec.UseNumber()
	decodeErr := dec.Decode(&obj)
	if decodeErr != nil {
		fmt.Printf("decode error: %v", decodeErr)
		return obj, decodeErr
	}
	return obj, nil
}

// MsgCompress compress bytes to s
func MsgCompress(msg TranData) (s string, err error) {
	var b []byte
	if b, err = json.Marshal(msg); err == nil {
		var buf bytes.Buffer
		if w, err := zlib.NewWriterLevel(&buf, zlib.BestCompression); err == nil {
			w.Write(b)
			w.Close()
			s = base64.URLEncoding.EncodeToString(b)
		}
	}
	return
}

func (w *WSClient) msgProcess(msg TranData) {
	log.Debugf("%v", msg)
	jgc := msg.RawMsg
	var gc gin.Context
	json.Unmarshal(jgc, &gc)
	log.Debugln(gc.Request)
	//w.SendMessage(TranData{
	//UUID:            w.sn,
	//Type:            msg.Type,
	//Data:            nil,
	//RawMsg: []byte{},
	//})
}

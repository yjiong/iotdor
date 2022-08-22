package api

import (
	"bytes"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/http/pprof"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var subServerPort = "8888"

// APIserver ....
func APIserver(ma ManageAPI, port string) {
	//log.SetReportCaller(true)
	//log.SetFormatter(&log.TextFormatter{
	//ForceColors:     true,
	//FullTimestamp:   true,
	//TimestampFormat: "2006/01/02 15:04:05",
	//})
	dtr := NewIotdorTran(ma, subServerPort)
	router := mux.NewRouter()
	router.PathPrefix("/api/login").HandlerFunc(dtr.rawlogin)
	//router.PathPrefix("/api/ws/transport").HandlerFunc(dtr.tranSport)
	configRoute(router.PathPrefix("/config").Subrouter(), dtr)
	gatewayRoute(router.PathPrefix("/gateway").Subrouter(), dtr)
	deviceRoute(router.PathPrefix("/device").Subrouter(), dtr)
	PprofGroup(router)
	router.PathPrefix("/").Handler(http.FileServer(wus))
	log.Infof("iotdor start apt server at port:%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

// IotdorTran data colloect meter Transport
type IotdorTran struct {
	IotdorHTTPPort   string
	Mu               *sync.Mutex
	Wsc              map[string]*websocket.Conn
	clist            map[string]string
	IotdorTrs        map[string]*http.Transport
	loginFailedCount map[string]int
	ManageAPI
}

// NewIotdorTran ...
func NewIotdorTran(m ManageAPI, p string) *IotdorTran {
	return &IotdorTran{
		IotdorHTTPPort:   p,
		Mu:               new(sync.Mutex),
		Wsc:              make(map[string]*websocket.Conn),
		clist:            make(map[string]string),
		loginFailedCount: make(map[string]int),
		IotdorTrs:        make(map[string]*http.Transport),
		ManageAPI:        m,
	}
}

// SetTR set http.Transport from websocket conn
func (dtr *IotdorTran) SetTR(sn string, wsc *websocket.Conn) {
	dtr.Mu.Lock()
	defer dtr.Mu.Unlock()
	dtr.Wsc[sn] = wsc
	dtr.IotdorTrs[sn] = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		Dial: func(nw, addr string) (net.Conn, error) {
			return wsc.UnderlyingConn(), nil
		},
	}
	dtr.clist[sn] = fmt.Sprintf("Last communication time: %s", time.Now().Format("2006-01-02 15:04:05"))
	log.Infof("Iotdor %s conn on connected %s\n", sn, wsc.RemoteAddr().String())
}

func (dtr *IotdorTran) iotdorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	je, _ := json.Marshal(dtr.clist)
	w.Write(je)
}

func (dtr *IotdorTran) request(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sn := params["Iotdorsn"]
	tr, ok := dtr.IotdorTrs[sn]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	rawurl := r.URL.String()
	IotdorURL := strings.TrimPrefix(rawurl, fmt.Sprintf("/%s", sn))
	rurl := fmt.Sprintf("http://localhost:%s%s", dtr.IotdorHTTPPort, IotdorURL)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if req, err := http.NewRequest(r.Method, rurl, nil); err == nil {
		req.Header = r.Header
		req.Body = r.Body
		resp, err := tr.RoundTrip(req)
		if err == nil {
			w.WriteHeader(resp.StatusCode)
			w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
			bf := bytes.Buffer{}
			bf.ReadFrom(resp.Body)
			w.Write(bf.Bytes())
			resp.Body.Close()
		} else {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Iotdor %s lost connect ...", sn)))
			log.Errorf("%s \n", err)
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

// RequestGin ....
func (dtr *IotdorTran) RequestGin(c *gin.Context) {
	dtr.Mu.Lock()
	defer dtr.Mu.Unlock()
	sn := c.Param("Iotdorsn")
	tr, ok := dtr.IotdorTrs[sn]
	if !ok {
		c.Status(http.StatusForbidden)
		return
	}
	ginurl := c.Request.URL.String()
	IotdorURL := strings.TrimPrefix(ginurl, fmt.Sprintf("/%s", sn))
	rurl := fmt.Sprintf("http://localhost:%s%s", dtr.IotdorHTTPPort, IotdorURL)
	if req, err := http.NewRequest(c.Request.Method, rurl, nil); err == nil {
		req.Header = c.Request.Header
		req.Body = c.Request.Body
		resp, err := tr.RoundTrip(req)
		if err == nil {
			c.Render(resp.StatusCode, render.Reader{
				ContentType:   resp.Header.Get("Content-Type"),
				ContentLength: resp.ContentLength,
				Reader:        resp.Body,
			})
		} else {
			log.Errorf("%+v %s \n", resp, err)
		}
		resp.Body.Close()
		return
	}
	c.Status(http.StatusInternalServerError)
}

func (dtr *IotdorTran) tranSport(w http.ResponseWriter, r *http.Request) {
	sn := r.Header.Get("sn")
	if sn == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	wconn, err := websocket.Upgrade(w, r, nil, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dtr.SetTR(sn, wconn)
}

//go:embed static/*
var wstatic embed.FS

type webui struct {
	wuEmbed embed.FS
	path    string
}

func (w *webui) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(w.path, filepath.FromSlash(path.Clean("/"+name)))
	file, err := w.wuEmbed.Open(fullName)
	wf := &WebuiFile{
		File: file,
	}
	return wf, err
}

// WebuiFile ......
type WebuiFile struct {
	io.Seeker
	fs.File
}

// Readdir ......
func (*WebuiFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

var wus = &webui{
	wuEmbed: wstatic,
	path:    "static",
}

func (dtr *IotdorTran) login(c *gin.Context) {
	dtr.rawlogin(c.Writer, c.Request)
}

func respError(code int, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{%q: %q}", "error", err.Error())))
}

func respJSON(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	jb, _ := json.Marshal(obj)
	w.Write(jb)
}

func (dtr *IotdorTran) delayForLogin(host string) {
	<-time.After(time.Minute * 3)
	delete(dtr.loginFailedCount, host)
}

func (dtr *IotdorTran) rawlogin(w http.ResponseWriter, req *http.Request) {
	var loginUser map[string]string
	if dtr.loginFailedCount[req.Host] > 3 {
		respError(200, w, errors.New("您错误登录太频繁了，休息一下吧"))
		return
	}
	err := decodeJSON(req, &loginUser)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	if err != nil {
		respError(200, w, err)
		return
	}
	euser, uerr := dtr.UserInfo(loginUser["username"])
	if uerr != nil || euser == nil {
		err = uerr
		if euser == nil {
			err = fmt.Errorf("user:%s not exist", loginUser["username"])
		}
		respError(200, w, err)
		dtr.loginFailedCount[req.Host]++
		go dtr.delayForLogin(req.Host)
		return
	}
	dpw, _ := hex.DecodeString(euser.Passwd)
	if derr := bcrypt.CompareHashAndPassword(dpw, []byte(loginUser["password"])); derr != nil {
		respError(200, w, errors.New("用户名或密码错误"))
		dtr.loginFailedCount[req.Host]++
		go dtr.delayForLogin(req.Host)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = float64(time.Now().Add(time.Hour * time.Duration(360)).Unix())
	claims["iat"] = float64(1)
	claims["uname"] = euser.Name
	token.Claims = claims
	dtr.UpdateUserLoginInfo(euser.Name, req.RemoteAddr)

	if tokenstr, err := token.SignedString(dpw); err == nil {
		tkm := (map[string]string{"Token": tokenstr})
		if loginUser["password"] == "123456" {
			tkm["msg"] = "尽快修改您的密码"
		}
		respJSON(w, tkm)
	} else {
		respError(http.StatusUnauthorized, w, err)
	}
}

func decodeJSON(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	decoder := json.NewDecoder(req.Body)
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}

// PprofGroup ....
func PprofGroup(r *mux.Router) {
	r.PathPrefix("/debug/pprof").HandlerFunc(pprof.Index)
	r.PathPrefix("/debug/heap").Handler(pprof.Handler("heap"))
	r.PathPrefix("/debug/goroutine").Handler(pprof.Handler("goroutine"))
	r.PathPrefix("/debug/allocs").Handler(pprof.Handler("allocs"))
	r.PathPrefix("/debug/block").Handler(pprof.Handler("block"))
	r.PathPrefix("/debug/threadcreate").Handler(pprof.Handler("threadcreate"))
	r.PathPrefix("/debug/cmdline").HandlerFunc(pprof.Cmdline)
	r.PathPrefix("/debug/profile").HandlerFunc(pprof.Profile)
	r.PathPrefix("/debug/symbol").HandlerFunc(pprof.Symbol)
	r.PathPrefix("/debug/symbol").HandlerFunc(pprof.Symbol).Methods("POST")
	r.PathPrefix("/debug/trace").HandlerFunc(pprof.Trace)
	r.PathPrefix("/debug/mutex").Handler(pprof.Handler("mutex"))
}

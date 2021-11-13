package main

import (
	"bytes"
	"embed"
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
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	//"github.com/gin-gonic/gin/render"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var serverPort = ":8899"
var subServerPort = "8888"

type controller struct {
}

// APIserver ....
func APIserver() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
	})
	dtr := NewiotdorTran(subServerPort)
	router := mux.NewRouter()
	router.PathPrefix("/api/login").HandlerFunc(dtr.rawlogin)
	router.PathPrefix("/api/ws/transport").HandlerFunc(dtr.tranSport)
	router.PathPrefix("/{iotdorsn}/iotdor/{any}").HandlerFunc(dtr.request).Methods("GET", "PUT", "POST", "DELETE")
	router.PathPrefix("/iotdors/list").HandlerFunc(dtr.iotdorList)
	PprofGroup(router)
	router.PathPrefix("/").Handler(http.FileServer(wus))
	http.ListenAndServe(":8899", router)
	/*
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		//ginpprof.Wrapper(router)
		//router.GET("/", gin.WrapH(http.FileServer(wus)))
			router.StaticFS("/broker", wus)
			dtr := NewiotdorTran(subServerPort)
			router.POST("/api/login", dtr.login)
			webAPI(router, dtr)
			err := router.Run(serverPort)
			if err != nil {
				fmt.Println("HTTP server failed ", err)
				os.Exit(1)
			}
	*/
}

// iotdorTran data colloect meter Transport
type iotdorTran struct {
	iotdorHTTPPort string
	Mu             *sync.Mutex
	Wsc            map[string]*websocket.Conn
	clist          map[string]string
	iotdorTrs      map[string]*http.Transport
}

// NewiotdorTran ...
func NewiotdorTran(p string) *iotdorTran {
	return &iotdorTran{
		iotdorHTTPPort: p,
		Mu:             new(sync.Mutex),
		Wsc:            make(map[string]*websocket.Conn),
		clist:          make(map[string]string),
		iotdorTrs:      make(map[string]*http.Transport),
	}
}

// SetTR set http.Transport from websocket conn
func (dtr *iotdorTran) SetTR(sn string, wsc *websocket.Conn) {
	dtr.Mu.Lock()
	defer dtr.Mu.Unlock()
	dtr.Wsc[sn] = wsc
	dtr.iotdorTrs[sn] = &http.Transport{
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
	log.Infof("iotdor %s conn on connected %s\n", sn, wsc.RemoteAddr().String())
}

func (dtr *iotdorTran) iotdorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	je, _ := json.Marshal(dtr.clist)
	w.Write(je)
}

func (dtr *iotdorTran) request(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sn := params["iotdorsn"]
	tr, ok := dtr.iotdorTrs[sn]
	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	rawurl := r.URL.String()
	iotdorURL := strings.TrimPrefix(rawurl, fmt.Sprintf("/%s", sn))
	rurl := fmt.Sprintf("http://localhost:%s%s", dtr.iotdorHTTPPort, iotdorURL)
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
			w.Write([]byte(fmt.Sprintf("iotdor %s lost connect ...", sn)))
			log.Errorf("%s \n", err)
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

// RequestGin ....
func (dtr *iotdorTran) RequestGin(c *gin.Context) {
	dtr.Mu.Lock()
	defer dtr.Mu.Unlock()
	sn := c.Param("iotdorsn")
	tr, ok := dtr.iotdorTrs[sn]
	if !ok {
		c.Status(http.StatusForbidden)
		return
	}
	ginurl := c.Request.URL.String()
	iotdorURL := strings.TrimPrefix(ginurl, fmt.Sprintf("/%s", sn))
	rurl := fmt.Sprintf("http://localhost:%s%s", dtr.iotdorHTTPPort, iotdorURL)
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

func webAPI(router *gin.Engine, dtr *iotdorTran) {
	//router.GET("/:iotdorsn/iotdor/*any", dtr.Request)
	//router.POST("/:iotdorsn/iotdor/*any", dtr.Request)
	//router.PUT("/:iotdorsn/iotdor/*any", dtr.Request)
}

func (dtr *iotdorTran) tranSport(w http.ResponseWriter, r *http.Request) {
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

var jwts = []byte("HJ212-yjiong@msn.com")

// TUser ...
type TUser struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	//IsAdmin  bool   `db:"is_admin" json:"is_admin"`
}

func (dtr *iotdorTran) login(c *gin.Context) {
	dtr.rawlogin(c.Writer, c.Request)
}

func validateToken(c *gin.Context) {
	validate(c.Writer, c.Request, func(w http.ResponseWriter, r *http.Request) { c.Next() })
	return
}

func validate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return jwts, nil
		}); err == nil {
		if token.Valid {
			next(w, r)
			return
		}
	}
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	w.WriteHeader(http.StatusUnauthorized)
	je, _ := json.Marshal(map[string]string{"error": "登录失败，或session timeout"})
	w.Write(je)
	return
}

func (dtr *iotdorTran) rawlogin(w http.ResponseWriter, req *http.Request) {
	var loginUser TUser
	err := decodeJSON(req, &loginUser)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	if err != nil {
		je, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(je)
		return
	}
	if !(loginUser.Username == "admin" && loginUser.Password == "lciotdor") {
		err = errors.Errorf("username not exist or password error")
		w.WriteHeader(http.StatusNotAcceptable)
		je, _ := json.Marshal(map[string]string{"error": "username not exist or password error"})
		w.Write(je)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = float64(time.Now().Add(time.Hour * time.Duration(360)).Unix())
	claims["iat"] = float64(1)
	token.Claims = claims
	if tokenstr, err := token.SignedString(jwts); err == nil {
		jr, _ := json.Marshal(map[string]string{"Token": tokenstr})
		w.Write(jr)
	} else {
		je, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(je)
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

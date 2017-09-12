package router

import (
	"net"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const (
	//AppName 端口
	AppName = "go-dbm"
	//Concurrency Concurrency
	Concurrency = 100000
	//DisableKeepalive DisableKeepalive
	DisableKeepalive = true
	//Timeout Timeout
	Timeout = 60 * time.Second
	//MaxConnsPerIP MaxConnsPerIP
	MaxConnsPerIP = 100000
	//MaxRequestsPerConn MaxRequestsPerConn
	MaxRequestsPerConn = 100000
	//MaxKeepaliveDuration MaxKeepaliveDuration
	MaxKeepaliveDuration = 120 * time.Second
	//MaxRequestBodySize MaxRequestBodySize
	MaxRequestBodySize = 512 * 1024 * 1024
)

//APIRoute 路由
type APIRoute struct {
	method []string
	handle func(*fasthttp.RequestCtx)
}

//Decorator 析构
type Decorator struct {
	RunFuc  func(*fasthttp.RequestCtx)
	PathStr string
}

//Decorator 判断 ip
func (d Decorator) Decorator(ctx *fasthttp.RequestCtx) {
	d.RunFuc(ctx)
	return
}

var dct Decorator

func getRouter() *fasthttprouter.Router {
	bs := NewBaseServer()
	APIRouteList := map[string]APIRoute{
		"/":       {[]string{"GET", "POST"}, bs.HandleRoot},
		"/status": {[]string{"GET"}, bs.GetStatus},
	}
	//app router

	router := &fasthttprouter.Router{
		RedirectTrailingSlash:  true,
		RedirectFixedPath:      true,
		HandleMethodNotAllowed: true,
		HandleOPTIONS:          true,
	}

	for k, v := range APIRouteList {
		for _, value := range v.method {
			dct.PathStr = k
			dct.RunFuc = v.handle
			router.Handle(value, k, dct.Decorator)
		}
	}
	return router
}

//GetServer 获取服务
func GetServer() *fasthttp.Server {
	r := getRouter()
	if r != nil {
		return &fasthttp.Server{
			Handler:              r.Handler,
			Name:                 AppName,
			Concurrency:          Concurrency,
			DisableKeepalive:     DisableKeepalive,
			ReadTimeout:          Timeout,
			WriteTimeout:         Timeout,
			MaxConnsPerIP:        MaxConnsPerIP,
			MaxRequestsPerConn:   MaxRequestsPerConn,
			MaxKeepaliveDuration: MaxKeepaliveDuration,
			MaxRequestBodySize:   MaxRequestBodySize,
		}
	}
	return nil
}

//Run 运行
func Run(n net.Listener) {
	server := GetServer()
	if server != nil {
		if err := server.Serve(n); err != nil {
			panic(err)
		}
	} else {
		panic("server is nil")
	}
}

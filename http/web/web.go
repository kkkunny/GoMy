package web

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// 中间件
type Middleware struct {
	name    string            // 名字
	handler MiddleHandlerFunc // 处理函数
}

// 新建一个默认配置的服务器
func NewDefault() *Web {
	serveMux := NewServerMux()
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serveMux,
	}
	web := &Web{server: server, serveMux: serveMux}
	web.AddMiddleware("ReqLog", MidRequestLog)
	return web
}

// 新建一个服务器
func New(addr string) *Web {
	serveMux := NewServerMux()
	server := &http.Server{
		Addr:    addr,
		Handler: serveMux,
	}
	web := &Web{server: server, serveMux: serveMux}
	web.AddMiddleware("ReqLog", MidRequestLog)
	return web
}

// 服务器
type Web struct {
	server      *http.Server // 服务器
	serveMux    *ServerMux   // 多路复用器
	middlewares []Middleware // 中间件列表
}

// 路由
func (this *Web) Route(name string, methods []string, pattern string, handler HandlerFunc) {
	// 中间件
	for i := len(this.middlewares) - 1; i >= 0; i-- {
		handler = this.middlewares[i].handler(handler)
	}
	// 路由
	if err := this.serveMux.Handle(name, methods, pattern, handler); err != nil {
		panic(err)
	}
}

// 增加中间件
func (this *Web) AddMiddleware(name string, handler MiddleHandlerFunc) {
	// 错误
	if handler == nil {
		panic("web: this middleware's handlerfunc is nil")
	}
	for _, v := range this.middlewares {
		if v.name == name {
			panic("web: this middleware:" + name + " is exist")
		}
	}
	// 增加
	mid := Middleware{name: name, handler: handler}
	this.middlewares = append(this.middlewares, mid)
}

// 静态文件夹
func (this *Web) Static(path string) {
	// 获取与文件名
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fileName := info.Name()
	pro := fmt.Sprintf("/%s/", fileName) // 前缀
	// 文件服务
	fileServer := http.FileServer(http.Dir(path))
	handle := HandlerFunc(func(ctx *Context) error {
		if p := strings.TrimPrefix(ctx.GetUrl().Path, pro); len(p) < len(ctx.GetUrl().Path) {
			r2 := new(http.Request)
			*r2 = *ctx.req
			r2.URL = new(url.URL)
			*r2.URL = *ctx.GetUrl()
			r2.URL.Path = p
			fileServer.ServeHTTP(ctx.writer, r2)
		}
		return nil
	})
	// 路由
	if err := this.serveMux.Handle(fileName, []string{http.MethodGet}, pro, handle); err != nil {
		panic(err)
	}
}

// 模板文件夹
func (this *Web) Templates(path string) {
	if err := this.serveMux.SetTemplatesFolder(path); err != nil {
		panic(err)
	}
}

// 白名单
func (this *Web) AllowIp(ip ...string) {
	for _, i := range ip {
		this.serveMux.whiteIpMaps[i] = 0
	}
}

// 黑名单
func (this *Web) NotAllowIp(ip ...string) {
	for _, i := range ip {
		this.serveMux.blackIpMaps[i] = 0
	}
}

// 开始监听
func (this *Web) Run() {
	_ = this.serveMux.Log.WriteInfoLog("Web server starts running...")
	_ = this.serveMux.Log.WriteInfoLog("listen on: http://" + this.server.Addr + "/")
	if err := this.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

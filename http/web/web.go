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
func NewDefault(addr ...string) *Web {
	serveMux := NewServerMux()
	var address string
	if len(addr) == 0{
		address = "127.0.0.1:8080"
	}else{
		address = addr[0]
	}
	server := &http.Server{
		Addr:    address,
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
	return &Web{server: server, serveMux: serveMux}
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
	handle := HandlerFunc(func(ctx *Context) {
		if p := strings.TrimPrefix(ctx.GetUrl().Path, pro); len(p) < len(ctx.GetUrl().Path) {
			r2 := new(http.Request)
			*r2 = *ctx.req
			r2.URL = new(url.URL)
			*r2.URL = *ctx.GetUrl()
			r2.URL.Path = p
			fileServer.ServeHTTP(ctx.writer, r2)
		}
	})
	// 路由
	if err := this.serveMux.Handle(fileName, []string{http.MethodGet}, pro, handle); err != nil {
		panic(err)
	}
}

// 开始监听
func (this *Web) Run() {
	fmt.Println("web: Web server starts running......")
	fmt.Println("web: address:http://" + this.server.Addr + "/")
	fmt.Println()
	if err := this.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

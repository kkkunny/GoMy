package web

import "fmt"

// 请求日志
func MidRequestLog(handler HandlerFunc) HandlerFunc {
	return func(ctx *Context) error {
		ip := ctx.GetRequestIp()
		fmt.Printf("request: a new [%s] request from IP: %s to ROUTE: %s\n接收数据:[%s]\n", ctx.GetMethod(), ip, ctx.GetUrl(), string(ctx.GetData()))
		return handler(ctx)
	}
}

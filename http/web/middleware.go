package web

import "fmt"

// 请求日志
func MidRequestLog(handler HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		ip := ctx.GetRequestIp()
		if len(ip) > 0 {
			fmt.Println(fmt.Sprintf("request: a new [%s] request from IP: %s to ROUTE: %s\n\t接收数据:[%s]", ctx.GetMethod(), ip, ctx.GetUrl(), string(ctx.GetData())))
		}
		handler(ctx)
	}
}

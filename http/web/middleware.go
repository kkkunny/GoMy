package web

import "fmt"

// 请求日志
func MidRequestLog(handler HandlerFunc) HandlerFunc {
	return func(ctx *Context) {
		ip := ctx.GetRequestIp()
		if len(ip) > 0 {
			fmt.Println(fmt.Sprintf("request: A new [%s] request from IP: %s", ctx.GetMethod(), ip))
		}
		handler(ctx)
	}
}

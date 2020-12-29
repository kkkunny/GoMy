package web

import "fmt"

// 请求日志
func MidRequestLog(handler HandlerFunc)HandlerFunc{
	return func(ctx *Context){
		ip := ctx.GetRequestIp()
		if len(ip) > 0{
			fmt.Println("request: A new request from IP: " + ip)
		}
		handler(ctx)
	}
}
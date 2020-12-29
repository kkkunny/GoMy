package web

// 处理器接口
type Handler interface {
	Handle(*Context) // 处理
}

// 处理器函数
type HandlerFunc func(*Context)
func (this HandlerFunc) Handle(r *Context){
	this(r)
}

// 中间件函数
type MiddleHandlerFunc func(HandlerFunc)HandlerFunc
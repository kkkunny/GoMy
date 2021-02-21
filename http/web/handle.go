package web

// 处理器接口
type Handler interface {
	Handle(*Context) error // 处理
}

// 处理器函数
type HandlerFunc func(*Context) error

func (this HandlerFunc) Handle(r *Context) error {
	return this(r)
}

// 中间件函数
type MiddleHandlerFunc func(HandlerFunc) HandlerFunc

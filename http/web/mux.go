package web

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

// 新建一个多路复用器
func NewServerMux() *ServerMux {
	mux := &ServerMux{
		tree:        NewRouteTree(),
		content:     make(map[string]*route),
		WhiteIpMaps: make(map[string]byte),
		BlackIpMaps: make(map[string]byte),
	}
	// url反射
	templateFuncs["url"] = func(name string) (string, error) {
		return urlReflex(mux.content, name)
	}
	return mux
}

// 多路复用器
type ServerMux struct {
	tree            *routeTree        // 路由树
	content         map[string]*route // 内容
	templatesFolder string            // 模板文件夹
	WhiteIpMaps     map[string]byte   // 白名单
	BlackIpMaps     map[string]byte   // 黑名单
}

// 检查请求方式是否被允许
func (this *ServerMux) isMethodAllowed(method string, r *route) bool {
	if len(r.methods) == 0 {
		return true
	}
	for _, v := range r.methods {
		if method == v {
			return true
		}
	}
	return false
}

// 是否允许ip访问
func (this *ServerMux) isIpAllowed(ip string) bool {
	if len(this.WhiteIpMaps) > 0 {
		if _, ok := this.WhiteIpMaps[ip]; ok {
			return true
		}
		return false
	}
	if len(this.BlackIpMaps) > 0 {
		if _, ok := this.BlackIpMaps[ip]; ok {
			return false
		}
		return true
	}
	return true
}

// 处理请求方式设置
func (this *ServerMux) handleMethods(r *route) error {
	// 任意请求方式
	if len(r.methods) != 0 {
		for _, v := range r.methods {
			switch v {
			case "ANY":
				r.methods = []string{}
				return nil
			case http.MethodGet:
				continue
			case http.MethodPost:
				continue
			case http.MethodPut:
				continue
			case http.MethodHead:
				continue
			case http.MethodPatch:
				continue
			case http.MethodTrace:
				continue
			case http.MethodOptions:
				continue
			case http.MethodDelete:
				continue
			case http.MethodConnect:
				continue
			default:
				return errors.New("web: wrong request methods is setted")
			}
		}
	}
	return nil
}

// 设置模板文件夹
func (this *ServerMux) SetTemplatesFolder(path string) error {
	// 确认是否存在
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	if path[len(path)-1] != '/' {
		path = path + "/"
	}
	this.templatesFolder = path
	return nil
}

// 创建路由
func (this *ServerMux) Handle(name string, methods []string, pattern string, handler HandlerFunc) error {
	// 错误
	if _, ok := this.content[name]; ok {
		return errors.New("web: the name '" + name + "' exist")
	}
	if len(pattern) <= 0 || pattern[0] != '/' {
		return errors.New("web: the pattern:" + pattern + " is err")
	}
	if handler == nil {
		return errors.New("web: the handler is nil")
	}
	r := &route{name: name, methods: methods, pattern: pattern, handler: handler}
	// 处理请求方式
	if err := this.handleMethods(r); err != nil {
		return err
	}
	// 创建路由节点
	err := this.tree.CreateRoute(pattern, r)
	if err != nil {
		return err
	}
	this.content[name] = r
	return nil
}

// 路由
func (this *ServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 查找路由节点
	rt, params, err := this.tree.SearchRouteNode(r.URL.Path)
	con := NewContext(w, r, params, this)
	// 是否被放入名单
	if !this.isIpAllowed(con.GetRequestIp()) {
		return
	}
	// 没有路由节点
	if err != nil || rt == nil || rt.handler == nil {
		_ = con.ReturnNotFound()
		return
	}
	// 检查请求方式是否被允许
	if ok := this.isMethodAllowed(r.Method, rt); !ok {
		_ = con.ReturnMethodNotAllowed()
		return
	}
	// 调用路由函数
	if err = rt.handler.Handle(con); err != nil {
		fmt.Println("web server error: " + err.Error())
	}
}

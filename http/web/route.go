package web

import (
	"errors"
	"strings"
)

// 路由
type route struct {
	name    string      // 名字
	methods []string    // 允许的方法
	pattern string      // url
	handler HandlerFunc // 处理器
}

// 获取url(反射用)
func (this *route) GetUrl() string {
	index := strings.LastIndex(this.pattern, ":")
	if index >= 0 {
		return this.pattern[:index]
	}
	return this.pattern
}

// 路由树节点
type routeTreeNode struct {
	father    *routeTreeNode            // 父节点
	sons      map[string]*routeTreeNode // 子节点
	route     *route                    // 路由
	back      bool                      // 是否允许回溯
	any       bool                      // 是有是带参路由
	paramName string                    // 参数名
}

// 新建路由树
func NewRouteTree() *routeTree {
	return &routeTree{root: &routeTreeNode{sons: make(map[string]*routeTreeNode), route: nil, back: false}}
}

// 路由树
type routeTree struct {
	root *routeTreeNode // 根节点
}

// 创建节点
func (this *routeTree) CreateRoute(url string, r *route) error {
	// 分离url
	sps := strings.Split(url, "/")[1:]
	cursor := this.root
	for k, sp := range sps {
		if sp != "" && sp[0] == ':' { // 带参路由
			son, ok := cursor.sons[":"]
			if !ok {
				node := &routeTreeNode{father: cursor, sons: make(map[string]*routeTreeNode), route: nil, back: false, any: true, paramName: sp[1:]}
				cursor.sons[":"] = node
				cursor = node
			} else {
				cursor = son
			}
		} else if sp != "" { // 如果不为空
			son, ok := cursor.sons[sp]
			if !ok {
				node := &routeTreeNode{father: cursor, sons: make(map[string]*routeTreeNode), route: nil, back: false}
				cursor.sons[sp] = node
				cursor = node
			} else {
				cursor = son
			}
		} else { // 为空
			if k == len(sps)-1 { // 最后一个为空时允许回溯
				cursor.back = true
			} else {
				return errors.New("routetree: the url:" + url + " is error")
			}
		}
	}
	if cursor.route != nil {
		return errors.New("routetree: this route is exist")
	} else {
		cursor.route = r
	}
	return nil
}

// 寻找路由节点，若无则返回上级节点和错误
func (this *routeTree) SearchRouteNode(url string) (*route, map[string]string, error) {
	// 路由参数
	params := make(map[string]string)
	// 分离url
	sps := strings.Split(url, "/")[1:]
	// 寻找节点
	cursor := this.root
	for _, sp := range sps {
		son, ok := cursor.sons[sp]
		if !ok {
			if sp != "" {
				paramson, paramok := cursor.sons[":"]
				if paramok { // 允许带参
					params[paramson.paramName] = sp
					cursor = paramson
					continue
				}
			}
			if cursor.back { // 是否允许回溯
				return cursor.route, params, nil
			} else {
				return nil, params, errors.New("routetree: no this route node")
			}
		} else {
			cursor = son
		}
	}
	return cursor.route, params, nil
}

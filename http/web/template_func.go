package web

import (
	"errors"
	"html/template"
)

// 模板函数定义
var templateFuncs = template.FuncMap{
	"add": add, "red": red, "mul": mul, "div": div, "mor": mor,
}

// add 加
func add(v1, v2 int) int {
	return v1 + v2
}

// red 减
func red(v1, v2 int) int {
	return v1 - v2
}

// mul 乘
func mul(v1, v2 int) int {
	return v1 * v2
}

// div 除
func div(v1, v2 int) int {
	return v1 / v2
}

// mor 余
func mor(v1, v2 int) int {
	return v1 % v2
}

// url反射
func urlReflex(routes map[string]*route, name string) (string, error) {
	if r, ok := routes[name]; ok {
		return r.GetUrl(), nil
	}
	return "", errors.New("do not found route '" + name + "'")
}

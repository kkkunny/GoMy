package re

import "regexp"

// 找到所有正则
func FindAll(com string, src string)[][]string{
	rg := regexp.MustCompile(com)
	return rg.FindAllStringSubmatch(src, -1)
}
// 替换所有正则
func ReplaceAll(com string, src string, repl string)string{
	rg := regexp.MustCompile(com)
	return rg.ReplaceAllString(src, repl)
}
// 是否存在匹配
func IsExist(com string, src string)bool{
	results := FindAll(com, src)
	if len(results) > 0{
		return true
	}
	return false
}
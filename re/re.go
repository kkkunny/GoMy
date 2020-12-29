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
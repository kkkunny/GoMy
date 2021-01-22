package path

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 获取父目录
func GetParentDirectory(path string) string {
	runes := []rune(path)
	last := strings.LastIndex(path, "\\")
	return string(runes[:last])
}

// 获取执行路径
func GetRunDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// 获取最后一级目录/文件
func GetLastPath(path string) string {
	runes := []rune(path)
	last := strings.LastIndex(path, "\\")
	if last < len(runes) {
		return string(runes[last+1:])
	} else {
		return ""
	}
}

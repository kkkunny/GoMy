package file

import (
	"encoding/hex"
	"os"
)

// 文件头与文件类型对照表
var fileTypeMap = map[string]string{
	"89504e470d0a1a0a0000": "png",
	"ffd8ffe000104a464946": "jpg",
	"47494638396126026f01": "gif",
}

// 获取文件类型
func GetFileType(file *os.File) (ft string, err error) {
	// 获取文件头
	_, err = file.Seek(0, 0) // 偏移到文件头
	if err != nil {
		return
	}
	buf := make([]byte, 10)
	if _, err = file.Read(buf); err != nil {
		return
	}
	defer file.Seek(0, 0) // 偏移到文件头
	filehead := hex.EncodeToString(buf)
	// 根据对照表获取类型
	if value, ok := fileTypeMap[filehead]; ok {
		return value, nil
	} else {
		return "unknown", nil
	}
}

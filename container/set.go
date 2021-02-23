package container

import (
	"fmt"
	"strings"
)

// 新建一个set
func NewSet() *Set {
	s := &Set{
		data: make(map[interface{}]byte),
	}
	return s
}

// set
type Set struct {
	data map[interface{}]byte
}

// 增加
func (this *Set) Add(elem interface{}) bool {
	if this.Contains(elem) {
		return false
	}
	this.data[elem] = 0
	return true
}

// 删除
func (this *Set) Remove(elem interface{}) {
	if this.Contains(elem) {
		delete(this.data, elem)
	}
}

// 包含
func (this *Set) Contains(elem interface{}) bool {
	if _, ok := this.data[elem]; ok {
		return true
	}
	return false
}

// 长度
func (this *Set) GetLength() int {
	return len(this.data)
}

// 清空
func (this *Set) Clear() {
	if this.GetLength() == 0 {
		return
	}
	for k, _ := range this.data {
		delete(this.data, k)
	}
}

// 交集
func (this *Set) Intersects(s *Set) {
	newData := make(map[interface{}]byte)
	for k, _ := range s.data {
		if this.Contains(k) {
			newData[k] = 0
		}
	}
	this.Clear()
	this.data = newData
}

// 并集
func (this *Set) Union(s *Set) {
	for k, _ := range s.data {
		this.Add(k)
	}
}

// 差集
func (this *Set) Differed(s *Set) {
	newData := make(map[interface{}]byte)
	for k, _ := range s.data {
		if !this.Contains(k) {
			newData[k] = 0
		}
	}
	for k, _ := range this.data {
		if !s.Contains(k) {
			newData[k] = 0
		}
	}
	this.Clear()
	this.data = newData
}

// 字符串
func (this *Set) GetString() string {
	var result []string
	for k, _ := range this.data {
		result = append(result, fmt.Sprintf("%v", k))
	}
	return "{" + strings.Join(result, ", ") + "}"
}

// 获取一个值
func (this *Set) Get() interface{} {
	for k, _ := range this.data {
		return k
	}
	return nil
}

// 转换成切片
func (this *Set) ToArray() []interface{} {
	var result []interface{}
	for k, _ := range this.data {
		result = append(result, k)
	}
	return result
}

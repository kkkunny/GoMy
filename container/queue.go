package container

import (
	"sync"
)

// 新建一个队列
func NewQueue() *Queue {
	return &Queue{lock: &sync.RWMutex{}, content: NewLinkList()}
}

// 队列
type Queue struct {
	lock    *sync.RWMutex
	content *LinkList
}

// 转化为文本
func (this *Queue) ToString() string {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.content.ToString()
}

// 获取值
func (this *Queue) Get() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.content.GetLength() > 0 {
		elem := this.content.Get(0)
		this.content.RemoveByIndex(0)
		return elem
	}
	return nil
}

// 放入值
func (this *Queue) Put(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.content.Append(value)
}

// 获取长度
func (this *Queue) GetLength() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.content.GetLength()
}

// 是否有该值
func (this *Queue) IsExist(value interface{}) bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.content.IsExist(value)
}

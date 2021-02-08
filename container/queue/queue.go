package queue

import (
	"github.com/kkkunny/GoMy/container/linklist"
	"sync"
)

// 新建一个队列
func New() *Queue {
	return &Queue{lock: &sync.Mutex{}, content: linklist.New()}
}

// 队列
type Queue struct {
	lock    *sync.Mutex
	content *linklist.LinkList
}

// 转化为文本
func (this *Queue) ToString() string {
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
	return this.content.GetLength()
}

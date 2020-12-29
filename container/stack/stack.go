package stack

import (
	"github.com/kkkunny/GoMy/container/linklist"
	"sync"
)

func New()*Stack{
	return &Stack{lock: sync.Mutex{}, content: linklist.New()}
}
// 栈
type Stack struct {
	lock sync.Mutex
	content *linklist.LinkList
}
// 入栈
func (this *Stack) Push(elem interface{}){
	this.lock.Lock()
	defer this.lock.Unlock()
	this.content.Append(elem)
}
// 出栈
func (this *Stack) Get()interface{}{
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.content.GetLength() > 0{
		elem := this.content.Get(this.content.GetLength()-1)
		this.content.RemoveByIndex(this.content.GetLength()-1)
		return elem
	}else{
		return nil
	}
}
// 获取长度
func (this *Stack) GetLength()int{
	return this.content.GetLength()
}
// 获取字符串
func (this *Stack) GetString()string{
	return this.content.GetString()
}
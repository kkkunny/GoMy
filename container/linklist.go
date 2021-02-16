package container

import (
	"fmt"
	"strings"
)

// 节点
type node struct {
	elem interface{} // 内容
	prev *node       // 上一个节点
	next *node       // 下一个节点
}

// 新建一个链表
func NewLinkList() *LinkList {
	return &LinkList{}
}

// 链表
type LinkList struct {
	head, end *node // 头，尾节点
	length    int   // 长度
}

// 根据下标寻找节点
func (this *LinkList) searchNodeByIndex(index int) *node {
	if index < 0 || index >= this.length {
		panic("Out of subscript range")
	} else {
		cursor, i := this.head, 0
		for cursor != nil {
			if i == index {
				break
			}
			i++
			cursor = cursor.next
		}
		return cursor
	}
}

// 根据元素寻找节点
func (this *LinkList) searchNodeByElem(elem interface{}) *node {
	cursor := this.head
	for cursor != nil {
		if cursor.elem == elem {
			return cursor
		}
		cursor = cursor.next
	}
	return nil
}

// 增加元素
func (this *LinkList) Add(index int, elem interface{}) {
	var cursor *node
	if index != this.length {
		cursor = this.searchNodeByIndex(index)
	}
	node := &node{elem: elem}
	if this.length == 0 { // 空链表
		this.head, this.end = node, node
	} else if cursor == this.head { // 头节点
		node.next = this.head
		this.head.prev = node
		this.head = node
	} else if cursor == nil { // 末尾
		this.end.next = node
		node.prev = this.end
		this.end = node
	} else {
		node.next = cursor
		node.prev = cursor.prev
		cursor.prev.next = node
		cursor.prev = node
	}
	this.length++
}

// 增加元素到末尾
func (this *LinkList) Append(elem interface{}) {
	this.Add(this.length, elem)
}

// 是否在链表中
func (this *LinkList) IsExist(elem interface{}) bool {
	if node := this.searchNodeByElem(elem); node != nil {
		return true
	}
	return false
}

// 获取下标
func (this *LinkList) GetIndex(elem interface{}) int {
	cursor, i := this.head, 0
	for cursor != nil {
		if cursor.elem == elem {
			return i
		}
		i++
		cursor = cursor.next
	}
	return -1
}

// 获取元素
func (this *LinkList) Get(index int) interface{} {
	node := this.searchNodeByIndex(index)
	return node.elem
}

// 删除节点
func (this *LinkList) removeNode(n *node) {
	if n == this.head { // 头节点
		this.head = n.next
		if this.head != nil {
			this.head.prev = nil
		}
	} else if n == this.end { // 尾节点
		this.end = n.prev
		if this.end != nil {
			this.end.next = nil
		}
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
	this.length--
	n.prev, n.next = nil, nil
}

// 删除下标
func (this *LinkList) RemoveByIndex(index int) interface{} {
	node := this.searchNodeByIndex(index)
	this.removeNode(node)
	return node.elem
}

// 删除第一个元素
func (this *LinkList) Remove(elem interface{}) bool {
	node := this.searchNodeByElem(elem)
	if node == nil {
		return false
	} else {
		this.removeNode(node)
		return true
	}
}

// 删除所有元素
func (this *LinkList) RemoveAll(elem interface{}) int {
	var num int
	for {
		node := this.searchNodeByElem(elem)
		if node == nil {
			break
		} else {
			this.removeNode(node)
			num++
		}
	}
	return num
}

// 获取长度
func (this *LinkList) GetLength() int {
	return this.length
}

// 转换成字符串
func (this *LinkList) ToString() string {
	var strs []string
	cursor := this.head
	for cursor != nil {
		strs = append(strs, fmt.Sprintf("%v", cursor.elem))
		cursor = cursor.next
	}
	temp := strings.Join(strs, ", ")
	return "[" + temp + "]"
}

// 函数遍历
func (this *LinkList) ErgodicFunc(handle func(interface{})) {
	cursor := this.head
	for cursor != nil {
		handle(cursor.elem)
		cursor = cursor.next
	}
}

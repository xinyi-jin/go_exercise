package main

import "fmt"

// 链表节点
type LinkNode struct {
	data interface{}
	next *LinkNode
}

// 链表的头结点和尾节点
type Link struct {
	head *LinkNode
	tail *LinkNode
}

// 头部插入
func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.head = node
		p.tail = p.head
		return
	}
	node.next = p.head
	p.head = node
}

// 尾部插入
func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.head == nil && p.tail == nil {
		p.head = node
		p.tail = p.head
		return
	}

	p.tail.next = node
	p.tail = node

}

// 遍历节点信息
func (p *Link) Traverse() {
	s := p.head

	for s != nil {
		fmt.Println("pNode info :", s.data)
		s = s.next
	}
}

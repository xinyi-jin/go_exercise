package stack

import "fmt"

// 链表栈
type LinkedListStack struct {
	Elem *LinkedList
}

type LinkedList struct {
	Val  int
	Next *LinkedList
}

func NewLinkedListSatck() *LinkedListStack {
	return &LinkedListStack{
		Elem: nil,
	}
}

func (ls *LinkedListStack) Push(elem int) {
	ls.Elem = &LinkedList{
		Val:  elem,
		Next: ls.Elem,
	}
}

func (ls *LinkedListStack) Pop() {
	if ls.Elem != nil {
		ls.Elem = ls.Elem.Next
	}
}

// 由于是指针，遍历输出的时候会改变源数据
func (ls *LinkedListStack) Traver() {
	for ls.Elem != nil {
		fmt.Print(ls.Elem.Val, " ")
		ls.Elem = ls.Elem.Next
	}
}

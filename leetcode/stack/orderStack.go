package stack

import (
	"fmt"
	"log"
)

// 顺序栈
type OrderStack struct {
	MaxElem int
	Top     int // default -1
	Elem    []int
}

func NewOrderStack(max int) *OrderStack {
	return &OrderStack{
		MaxElem: max,
		Top:     DEFAULT_TOP,
		Elem:    []int{},
	}
}

func (os *OrderStack) Push(elem int) {
	if os.Top >= os.MaxElem {
		log.Fatalln("OrderStack over stack!")
		return
	}
	os.Top++
	os.Elem = append(os.Elem, elem)
}

func (os *OrderStack) Pop() {
	if os.Top == -1 {
		log.Fatalln("OrderStack is nil!")
		return
	}
	os.Top--
	os.Elem = os.Elem[:len(os.Elem)-1]
}

func (os *OrderStack) Traver() {
	for i := os.Top; i >= 0; i-- {
		fmt.Printf("%v ", os.Elem[i])
	}
	fmt.Println()
}

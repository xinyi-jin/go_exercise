package alg1670

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Constructor()
	q.PushFront(1) // [1]
	q.PrintQueue()
	q.PushBack(2) // [1, 2]
	q.PrintQueue()
	q.PushMiddle(3) // [1, 3, 2]
	q.PrintQueue()
	q.PushMiddle(4) // [1, 4, 3, 2]
	q.PrintQueue()
	fmt.Print(q.PopFront(), " -> ") // 返回 1 -> [4, 3, 2]
	q.PrintQueue()
	fmt.Print(q.PopMiddle(), " -> ") // 返回 3 -> [4, 2]
	q.PrintQueue()
	fmt.Print(q.PopMiddle(), " -> ") // 返回 4 -> [2]
	q.PrintQueue()
	fmt.Print(q.PopBack(), " -> ") // 返回 2 -> []
	q.PrintQueue()
	fmt.Print(q.PopFront(), " -> ") // 返回 -1 -> [] （队列为空）
	q.PrintQueue()
}

func (fm *FrontMiddleBackQueue) PrintQueue() {
	fmt.Println(fm.queue)
}

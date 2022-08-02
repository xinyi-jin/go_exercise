package alg622

/* 622. 设计循环队列
设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。

循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。

你的实现应该支持如下操作：

MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。


示例：

MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
circularQueue.enQueue(1);  // 返回 true
circularQueue.enQueue(2);  // 返回 true
circularQueue.enQueue(3);  // 返回 true
circularQueue.enQueue(4);  // 返回 false，队列已满
circularQueue.Rear();  // 返回 3
circularQueue.isFull();  // 返回 true
circularQueue.deQueue();  // 返回 true
circularQueue.enQueue(4);  // 返回 true
circularQueue.Rear();  // 返回 4


提示：

所有的值都在 0 至 1000 的范围内；
操作数将在 1 至 1000 的范围内；
请不要使用内置的队列库。 */

// 思路：环形数组
type MyCircularQueue struct {
	head  int   // 头节点坐标
	tail  int   // 尾节点坐标
	queue []int // 数组
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k+1),
	}
}

func (mcq *MyCircularQueue) EnQueue(value int) bool {
	if mcq.IsFull() {
		return false
	}
	mcq.queue[mcq.tail] = value
	mcq.tail = (mcq.tail + 1) % len(mcq.queue)
	return true
}

func (mcq *MyCircularQueue) DeQueue() bool {
	if mcq.IsEmpty() {
		return false
	}
	mcq.head = (mcq.head + 1) % len(mcq.queue)
	return true
}

func (mcq *MyCircularQueue) Front() int {
	if mcq.IsEmpty() {
		return -1
	}
	return mcq.queue[mcq.head]
}

func (mcq *MyCircularQueue) Rear() int {
	if mcq.IsEmpty() {
		return -1
	}
	return mcq.queue[(mcq.tail-1+len(mcq.queue))%len(mcq.queue)]
}

func (mcq *MyCircularQueue) IsEmpty() bool {
	return mcq.tail == mcq.head
}

func (mcq *MyCircularQueue) IsFull() bool {
	return (mcq.tail+1)%len(mcq.queue) == mcq.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

//  不更新下标实际值解法 直接递增 进行模运算取值，不必考虑int类型最大值越界问题
/*  type MyCircularQueue struct {
	head  int   // 头节点坐标
	tail  int   // 尾节点坐标
	queue []int // 数组
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
	}
}

func (mcq *MyCircularQueue) EnQueue(value int) bool {
	if mcq.IsFull() {
		return false
	}
	mcq.queue[mcq.tail%len(mcq.queue)] = value
	mcq.tail++
	return true
}

func (mcq *MyCircularQueue) DeQueue() bool {
	if mcq.IsEmpty() {
		return false
	}
	mcq.head++
	return true
}

func (mcq *MyCircularQueue) Front() int {
	if mcq.IsEmpty() {
		return -1
	}
	return mcq.queue[mcq.head%len(mcq.queue)]
}

func (mcq *MyCircularQueue) Rear() int {
	if mcq.IsEmpty() {
		return -1
	}
	return mcq.queue[(mcq.tail-1)%len(mcq.queue)]
}

func (mcq *MyCircularQueue) IsEmpty() bool {
	return mcq.tail == mcq.head
}

func (mcq *MyCircularQueue) IsFull() bool {
	return mcq.tail-mcq.head == len(mcq.queue)
} */

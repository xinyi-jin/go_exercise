package alg641

/* 641. 设计循环双端队列
设计实现双端队列。

实现 MyCircularDeque 类:

MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。


示例 1：

输入
["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
[[3], [1], [2], [3], [4], [], [], [], [4], []]
输出
[null, true, true, true, false, 2, true, true, true, 4]

解释
MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4



提示：

1 <= k <= 1000
0 <= value <= 1000
insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty, isFull  调用次数不大于 2000 次 */

// 思路：环形数组
type MyCircularDeque struct {
	head  int   // 头节点坐标
	tail  int   // 尾节点坐标
	queue []int // 数组
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k+1),
	}
}

func (mcd *MyCircularDeque) InsertFront(value int) bool {
	if mcd.IsFull() {
		return false
	}
	mcd.head = (mcd.head - 1 + len(mcd.queue)) % len(mcd.queue)
	mcd.queue[mcd.head] = value
	return true
}

func (mcd *MyCircularDeque) InsertLast(value int) bool {
	if mcd.IsFull() {
		return false
	}
	mcd.queue[mcd.tail] = value
	mcd.tail = (mcd.tail + 1) % len(mcd.queue)
	return true
}

func (mcd *MyCircularDeque) DeleteFront() bool {
	if mcd.IsEmpty() {
		return false
	}
	mcd.head = (mcd.head + 1) % len(mcd.queue)
	return true
}

func (mcd *MyCircularDeque) DeleteLast() bool {
	if mcd.IsEmpty() {
		return false
	}
	mcd.tail = (mcd.tail - 1 + len(mcd.queue)) % len(mcd.queue)
	return true
}

func (mcd *MyCircularDeque) GetFront() int {
	if mcd.IsEmpty() {
		return -1
	}
	return mcd.queue[mcd.head]
}

func (mcd *MyCircularDeque) GetRear() int {
	if mcd.IsEmpty() {
		return -1
	}
	return mcd.queue[(mcd.tail-1+len(mcd.queue))%len(mcd.queue)]
}

func (mcd *MyCircularDeque) IsEmpty() bool {
	return mcd.tail == mcd.head
}

func (mcd *MyCircularDeque) IsFull() bool {
	return (mcd.tail+1)%len(mcd.queue) == mcd.head
}

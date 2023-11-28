package alg1670

type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (fm *FrontMiddleBackQueue) IsEmpty() bool {
	return len(fm.queue) == 0
}

func (fm *FrontMiddleBackQueue) PushFront(val int) {
	fm.queue = append([]int{val}, fm.queue...)
}

func (fm *FrontMiddleBackQueue) PushMiddle(val int) {
	idx := len(fm.queue) / 2
	// 临时切片保存插入位置后边元素
	temp := append([]int{}, fm.queue[idx:]...)
	fm.queue = append(fm.queue[:idx], val)
	fm.queue = append(fm.queue, temp...)
}

func (fm *FrontMiddleBackQueue) PushBack(val int) {
	fm.queue = append(fm.queue, val)
}

func (fm *FrontMiddleBackQueue) PopFront() int {
	if fm.IsEmpty() {
		return -1
	}
	res := fm.queue[0]
	fm.queue = fm.queue[1:]
	return res
}

func (fm *FrontMiddleBackQueue) PopMiddle() int {
	if fm.IsEmpty() {
		return -1
	}
	idx := len(fm.queue) / 2
	if len(fm.queue)%2 == 0 {
		idx--
	}
	res := fm.queue[idx]
	fm.queue = append(fm.queue[:idx], fm.queue[idx+1:]...)
	return res
}

func (fm *FrontMiddleBackQueue) PopBack() int {
	if fm.IsEmpty() {
		return -1
	}
	n := len(fm.queue)
	res := fm.queue[n-1]
	fm.queue = fm.queue[:n-1]
	return res
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */

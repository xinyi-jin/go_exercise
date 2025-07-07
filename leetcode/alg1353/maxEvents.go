package alg1353

import (
	"container/heap"
	"sort"
)

/* 1353. 最多可以参加的会议数目
中等
相关标签
premium lock icon
相关企业
提示
给你一个数组 events，其中 events[i] = [startDayi, endDayi] ，表示会议 i 开始于 startDayi ，结束于 endDayi 。

你可以在满足 startDayi <= d <= endDayi 中的任意一天 d 参加会议 i 。在任意一天 d 中只能参加一场会议。

请你返回你可以参加的 最大 会议数目。



示例 1：



输入：events = [[1,2],[2,3],[3,4]]
输出：3
解释：你可以参加所有的三个会议。
安排会议的一种方案如上图。
第 1 天参加第一个会议。
第 2 天参加第二个会议。
第 3 天参加第三个会议。
示例 2：

输入：events= [[1,2],[2,3],[3,4],[1,2]]
输出：4


提示：​​​​​​

1 <= events.length <= 105
events[i].length == 2
1 <= startDayi <= endDayi <= 105 */

// 思路：贪心算法 + 小根堆
func maxEvents(events [][]int) int {
	res := 0
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	maxDay := 0
	for _, v := range events {
		if v[1] > maxDay {
			maxDay = v[1]
		}
	}
	minHeap := &InitHeap{}
	heap.Init(minHeap)
	for i, j := 1, 0; i <= maxDay; i++ {
		for j < len(events) && events[j][0] <= i {
			heap.Push(minHeap, events[j][1])
			j++
		}
		for minHeap.Len() > 0 && (*minHeap)[0] < i {
			heap.Pop(minHeap)
		}
		if minHeap.Len() > 0 {
			heap.Pop(minHeap)
			res++
		}
	}
	return res
}

// 贪心算法 + 排序数组 提交会超时
func maxEventsEx(events [][]int) int {
	res := 0
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	maxDay := 0
	for _, v := range events {
		if v[1] > maxDay {
			maxDay = v[1]
		}
	}
	sortArr := InitHeap{}
	for i, j := 1, 0; i <= maxDay; i++ {
		for j < len(events) && events[j][0] <= i {
			sortArr = append(sortArr, events[j][1])
			j++
		}
		sort.Sort(sortArr)
		for len(sortArr) > 0 && sortArr[0] < i {
			sortArr = append(sortArr[:0], sortArr[1:]...)
		}
		if len(sortArr) > 0 {
			sortArr = append(sortArr[:0], sortArr[1:]...)
			res++
		}

	}
	return res
}

type InitHeap []int

func (a InitHeap) Len() int           { return len(a) }
func (a InitHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a InitHeap) Less(i, j int) bool { return a[i] < a[j] }
func (h *InitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *InitHeap) Push(a interface{}) {
	*h = append(*h, a.(int))
}

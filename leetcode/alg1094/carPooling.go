package alg1094

import "sort"

/* 1094. 拼车
提示
中等
305
相关企业
车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）

给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。

当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。



示例 1：

输入：trips = [[2,1,5],[3,3,7]], capacity = 4
输出：false
示例 2：

输入：trips = [[2,1,5],[3,3,7]], capacity = 5
输出：true


提示：

1 <= trips.length <= 1000
trips[i].length == 3
1 <= numPassengersi <= 100
0 <= fromi < toi <= 1000
1 <= capacity <= 105 */

// 思路：模拟 + 差分数组
// 1. 求每个站点位置上下人数
// 2. 校验当前站点是否超载（超载返回false）
func carPooling(trips [][]int, capacity int) bool {
	maxCap := 0
	for _, v := range trips {
		maxCap = max(maxCap, v[2])
	}
	diff := make([]int, maxCap+1)
	for _, v := range trips {
		diff[v[1]] += v[0]
		diff[v[2]] -= v[0]
	}
	num := 0
	for _, v := range diff {
		num += v
		if num > capacity {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type SortByEnter [][]int

func (a SortByEnter) Len() int           { return len(a) }
func (a SortByEnter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByEnter) Less(i, j int) bool { return a[i][1] < a[j][1] }

// 思路：模拟 + 排序 + 哈希
func carPoolingEx(trips [][]int, capacity int) bool {
	sort.Sort(SortByEnter(trips))
	num := 0
	leaveMap := make(map[int]int)
	startIdx := 1
	for _, v := range trips {
		num += v[0]
		leaveMap[v[2]] += v[0]
		for i := startIdx; i <= v[1]; i++ {
			if n, ok := leaveMap[i]; ok && n > 0 {
				delete(leaveMap, i)
				num -= n
			}
		}
		startIdx = v[1]
		if num > capacity {
			return false
		}
	}
	return true
}

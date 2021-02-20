package main

import "fmt"

/* 给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/degree-of-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */
type numMap struct {
	count      int //出现的最次数
	leftIndex  int //左索引
	rightIndex int //右索引
}

// 思路：使用哈希，记录每个数字出现的次数，左边索引和右边索引。满足相同大小度的子数组最小长度，应包含该数字全部值，即该数字第一次出现到最后一次出现，即是满足条件的最小子数组
// 遍历计算出得哈希表，即可得出结果
func findShortestSubArray(nums []int) int {
	numMaps := make(map[int]*numMap)
	for i, v := range nums {
		if m, ok := numMaps[v]; ok {
			m.count++
			m.rightIndex = i
			numMaps[v] = m
		} else {
			numMaps[v] = &numMap{1, i, i}
		}
	}
	maxCnt, res := 0, 0
	for _, v := range numMaps {
		if v.count > maxCnt {
			maxCnt, res = v.count, v.rightIndex-v.leftIndex+1
		} else if v.count == maxCnt {
			res = min(res, v.rightIndex-v.leftIndex+1)
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := make([]int, 10)
	nums = []int{1, 2, 2, 3, 1, 4, 2}
	result := findShortestSubArray(nums)
	fmt.Println("result:", result)

}

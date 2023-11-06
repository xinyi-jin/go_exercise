package alg421

import "sort"

/* 421. 数组中两个数的最大异或值
中等
583
相关企业
给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。



示例 1：

输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
示例 2：

输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
输出：127


提示：

1 <= nums.length <= 2 * 105
0 <= nums[i] <= 231 - 1 */

// 哈希+位运算
func findMaximumXOR(nums []int) (res int) {
	highestBit := 30
	for i := highestBit; i >= 0; i-- {
		tempMap := make(map[int]bool)
		for _, v := range nums {
			tempMap[v>>i] = true
		}

		next := res<<1 | 1
		found := false
		for _, v := range nums {
			if tempMap[v>>i^next] {
				found = true
				break
			}
		}
		if found {
			res = next
		} else {
			res = next - 1
		}

	}
	return
}

// 暴力解法(超时)
func findMaximumXOREx(nums []int) int {
	maxValue := 0
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			maxValue = max(maxValue, nums[i]^nums[j])
		}
	}

	return maxValue
}

// 暴力解法优化(仍超时)
func findMaximumXORExtion(nums []int) int {
	maxValue := 0
	sort.Ints(nums)
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i]+nums[j] < maxValue {
				break
			}
			maxValue = max(maxValue, nums[i]^nums[j])
		}
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

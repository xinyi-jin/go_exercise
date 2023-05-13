package alg2441

import "math"

/* 2441. 与对应负数同时存在的最大正整数
给你一个 不包含 任何零的整数数组 nums ，找出自身与对应的负数都在数组中存在的最大正整数 k 。

返回正整数 k ，如果不存在这样的整数，返回 -1 。



示例 1：

输入：nums = [-1,2,-3,3]
输出：3
解释：3 是数组中唯一一个满足题目要求的 k 。
示例 2：

输入：nums = [-1,10,6,7,-7,1]
输出：7
解释：数组中存在 1 和 7 对应的负数，7 的值更大。
示例 3：

输入：nums = [-10,8,6,7,-2,-3]
输出：-1
解释：不存在满足题目要求的 k ，返回 -1 。


提示：

1 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
nums[i] != 0 */

// 思路：直接计算 优先内存
func findMaxK(nums []int) int {
	res := -1
	for i, a := range nums[:len(nums)-1] {
		isMinus := a < 0
		for _, b := range nums[i+1:] {
			if isMinus != (b < 0) && a == -b {
				res = max(res, int(math.Abs(float64(a))))
				break
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 思路：直接计算 这种内存与时间
func findMaxKEx(nums []int) int {
	ans := -1
	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
		if abs(x) > ans && has[-x] {
			ans = abs(x)
		}

	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

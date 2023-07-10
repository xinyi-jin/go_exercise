package alg16

import (
	"math"
	"sort"
)

/* 16. 最接近的三数之和
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。



示例 1：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
示例 2：

输入：nums = [0,0,0], target = 1
输出：0


提示：

3 <= nums.length <= 1000
-1000 <= nums[i] <= 1000
-104 <= target <= 104 */

// 暴力解法
func threeSumClosest(nums []int, target int) int {
	res := math.MaxInt32
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for x := j + 1; x < n; x++ {
				sum := nums[i] + nums[j] + nums[x]
				if sum == target {
					return sum
				}
				if abs(sum-target) < abs(res-target) {
					res = sum
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

// 排序+双指针收缩
func threeSumClosestExtion(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	// 枚举 a
	for i := 0; i < n; i++ {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				// 如果和大于 target，移动 c 对应的指针
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

// 简单粗暴
func threeSumClosestEx(nums []int, target int) int {
	res := -1
	diff := 0
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for x := j + 1; x < n; x++ {
				tempDiff := 0
				sum := nums[i] + nums[j] + nums[x]
				if sum == target {
					return sum
				}
				if sum > target {
					tempDiff = sum - target

				} else {
					tempDiff = target - sum
				}
				if tempDiff <= diff || res == -1 {
					res = sum
					diff = tempDiff
				}
			}
		}
	}
	return res
}

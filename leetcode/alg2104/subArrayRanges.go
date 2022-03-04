package alg2104

import "sort"

/* 2104. 子数组范围和
给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。

返回 nums 中 所有 子数组范围的 和 。

子数组是数组中一个连续 非空 的元素序列。



示例 1：

输入：nums = [1,2,3]
输出：4
解释：nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0
[2]，范围 = 2 - 2 = 0
[3]，范围 = 3 - 3 = 0
[1,2]，范围 = 2 - 1 = 1
[2,3]，范围 = 3 - 2 = 1
[1,2,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4
示例 2：

输入：nums = [1,3,3]
输出：4
解释：nums 的 6 个子数组如下所示：
[1]，范围 = 最大 - 最小 = 1 - 1 = 0
[3]，范围 = 3 - 3 = 0
[3]，范围 = 3 - 3 = 0
[1,3]，范围 = 3 - 1 = 2
[3,3]，范围 = 3 - 3 = 0
[1,3,3]，范围 = 3 - 1 = 2
所有范围的和是 0 + 0 + 0 + 2 + 0 + 2 = 4
示例 3：

输入：nums = [4,-2,-3,4,1]
输出：59
解释：nums 中所有子数组范围的和是 59


提示：

1 <= nums.length <= 1000
-109 <= nums[i] <= 109


进阶：你可以设计一种时间复杂度为 O(n) 的解决方案吗？ */

// 思路:遍历
func subArrayRanges(nums []int) int64 {
	res := 0
	for i := 0; i < len(nums); i++ {
		maxNum, minNum := nums[i], nums[i]
		for _, v := range nums[i+1:] {
			minNum = min(minNum, v)
			maxNum = max(maxNum, v)
			res += maxNum - minNum
		}
	}
	return int64(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 暴力解 timeout 原因：底层排序耗时引起时间复杂度偏高
func subArrayRangesEx(nums []int) int64 {
	sum := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			temp := make([]int, j+1-i)
			copy(temp, nums[i:j+1])
			sort.Ints(temp)
			sum += temp[len(temp)-1] - temp[0]
		}
	}
	return int64(sum)
}

// 找所有的最大值与最小值 最优解。
func subArrayRangesExtion(nums []int) int64 {
	n := len(nums)
	left, right := 0, 0
	res := 0
	for i := 0; i < n; i++ {
		left, right = i, i
		for left-1 >= 0 && nums[left-1] < nums[i] {
			left--
		}
		for right+1 < n && nums[right+1] <= nums[i] {
			right++
		}
		res += (right - i + 1) * (i - left + 1) * nums[i] // 最大值

		left, right = i, i
		for left-1 >= 0 && nums[left-1] > nums[i] {
			left--
		}
		for right+1 < n && nums[right+1] >= nums[i] {
			right++
		}
		res -= (right - i + 1) * (i - left + 1) * nums[i] // 最小值
	}
	return int64(res)
}

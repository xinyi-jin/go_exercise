package alg1749

/* 1749. 任意子数组和的绝对值的最大值
给你一个整数数组 nums 。一个子数组 [numsl, numsl+1, ..., numsr-1, numsr] 的 和的绝对值 为 abs(numsl + numsl+1 + ... + numsr-1 + numsr) 。

请你找出 nums 中 和的绝对值 最大的任意子数组（可能为空），并返回该 最大值 。

abs(x) 定义如下：

如果 x 是负整数，那么 abs(x) = -x 。
如果 x 是非负整数，那么 abs(x) = x 。


示例 1：

输入：nums = [1,-3,2,3,-4]
输出：5
解释：子数组 [2,3] 和的绝对值最大，为 abs(2+3) = abs(5) = 5 。
示例 2：

输入：nums = [2,-5,1,-4,3,-2]
输出：8
解释：子数组 [-5,1,-4] 和的绝对值最大，为 abs(-5+1-4) = abs(-8) = 8 。


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104 */

/* 精辟：瞬间理解思路
yang-xiao-yue-3
杨晓月
来自新加坡
6 小时前
莽一看以为动态规划。草稿纸上画了画，想到求子串的和一般用前缀和，数组的前缀和的曲线会上下波动，那最大的子数组的和绝对值不就是前缀和曲线的波峰和波谷的差么。 */

/* 思路：由于子数组和等于两个前缀和的差，那么取前缀和中的最大值与最小值，它俩的差就是答案。

如果最大值在最小值右边，那么算的是最大子数组和。

如果最大值在最小值左边，那么算的是最小子数组和的绝对值（相反数）。 */
func maxAbsoluteSum(nums []int) int {
	var s, max, min int
	for _, v := range nums {
		s += v
		if s > max {
			max = s
		} else if s < min {
			min = s
		}
	}
	return max - min
}

// 思路：前缀和 暴力解法 力扣会超时
func maxAbsoluteSumEx(nums []int) int {
	max := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := 0
		for j := i; j < n; j++ {
			temp = temp + nums[j]
			if m := abs(temp); m > max {
				max = m
			}
		}
	}
	return max
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

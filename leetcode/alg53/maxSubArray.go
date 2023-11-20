package alg53

/* 53. 最大子数组和
中等
6.4K
相关企业
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。



示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104 */

// 动态规划
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

// 前缀和构建缓存(超时)
func maxSubArrayExtion(nums []int) int {
	res := nums[0]
	n := len(nums)
	cash := make([]int, n)
	cash[0] = nums[0]
	for i := 1; i < n; i++ {
		cash[i] = cash[i-1] + nums[i]
	}

	for i := 0; i < n; i++ {
		res = max(res, cash[i])
		for j := i + 1; j < n; j++ {
			res = max(res, cash[j]-cash[i])
		}
	}
	return res
}

// 直接遍历(超时)
func maxSubArrayEx(nums []int) int {
	res := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		sum := nums[i]
		res = max(res, sum)
		for j := i + 1; j < n; j++ {
			sum += nums[j]
			res = max(res, sum)
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

package main

import "fmt"

/* 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

// 198. 打家劫舍
// 思路：要么从零开始,遍历所有情况
func rob(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < l; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[l-1]
}

func robEx(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < l; i++ {
		first, second = second, max(first+nums[i], second)
	}

	return second
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func main() {
	// arr := []int{1, 2, 3, 1}
	arr := []int{2, 7, 9, 3, 1}
	// arr := []int{1, 2, 3, 18}

	fmt.Printf("res %d", rob(arr))
	fmt.Printf("res %d", robEx(arr))
}

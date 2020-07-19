package main

import "fmt"

/* 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

求所能获得硬币的最大数量。

说明:

你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
示例:

输入: [3,1,5,8]
输出: 167
解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/burst-balloons
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func main() {
	// [3,1,5,8]
	arrTest := []int{3, 1, 5, 8}
	result := maxCoins(arrTest)
	fmt.Println("result:", result)

}

// 思路：假设当前i是最后选择的气球，拆分成左右两部分
func maxCoins(nums []int) int {
	n := len(nums)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i < n+1; i++ {
		val[i] = nums[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for m := i + 1; m < j; m++ {
				sum := val[i] * val[m] * val[j]
				sum += dp[i][m] + dp[m][j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

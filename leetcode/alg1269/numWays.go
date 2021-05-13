package alg1269

/* 1269. 停在原地的方案数
有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。

每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。

给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。

由于答案可能会很大，请返回方案数 模 10^9 + 7 后的结果。 */

// 困难 官方题解
func numWays(steps, arrLen int) int {
	const mod = 1e9 + 7
	maxColumn := min(arrLen-1, steps)
	dp := make([]int, maxColumn+1)
	dp[0] = 1
	for i := 1; i <= steps; i++ {
		dpNext := make([]int, maxColumn+1)
		for j := 0; j <= maxColumn; j++ {
			dpNext[j] = dp[j]
			if j-1 >= 0 {
				dpNext[j] = (dpNext[j] + dp[j-1]) % mod
			}
			if j+1 <= maxColumn {
				dpNext[j] = (dpNext[j] + dp[j+1]) % mod
			}
		}
		dp = dpNext
	}
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

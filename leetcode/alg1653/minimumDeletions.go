package alg1653

/* 1653. 使字符串平衡的最少删除次数
中等
相关标签
premium lock icon
相关企业
提示
给你一个字符串 s ，它仅包含字符 'a' 和 'b'​​​​ 。

你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a' ，此时认为 s 是 平衡 的。

请你返回使 s 平衡 的 最少 删除次数。



示例 1：

输入：s = "aababbab"
输出：2
解释：你可以选择以下任意一种方案：
下标从 0 开始，删除第 2 和第 6 个字符（"aababbab" -> "aaabbb"），
下标从 0 开始，删除第 3 和第 6 个字符（"aababbab" -> "aabbbb"）。
示例 2：

输入：s = "bbaaaaabb"
输出：2
解释：唯一的最优解是删除最前面两个字符。


提示：

1 <= s.length <= 105
s[i] 要么是 'a' 要么是 'b'​ 。​ */

// 思路计算当前索引 前后需要删除的个数，取最小值
func minimumDeletions(s string) int {
	n := len(s)
	dp := make([][2]int, n+1)
	for i := 1; i <= n; i++ {
		// dp[i][0] 表示删除当前字符后，前面字符为a的个数
		// dp[i][1] 表示删除当前字符后，前面字符为b的个数
		if s[i-1] == 'a' {
			// 保留a，删除b
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][1], dp[i-1][0]) + 1
		} else { // s[i-1] == 'b'
			// 删除a，保留b
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][1], dp[i-1][0])
		}
	}

	return min(dp[n][0], dp[n][1])
}

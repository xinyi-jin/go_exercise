package main

import "fmt"

func main() {
	// s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
	// s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
	s1, s2, s3 := "aabcc", "dbbca", "aadbbbaccc"
	result := isInterleave(s1, s2, s3)
	fmt.Println("result", result)

}

// 使用dp思路解决
func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, s := len(s1), len(s2), len(s3)

	// 长度不匹配，直接返回false
	if m+n != s {
		return false
	}

	// 初始化
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			p := i + j - 1
			if i > 0 {
				dp[j] = dp[j] && s1[i-1] == s3[p]
			}
			if j > 0 {
				dp[j] = dp[j] || dp[j-1] && s2[j-1] == s3[p]
			}
		}
	}
	return dp[n]
}

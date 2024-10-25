package alg409

/* 409. 最长回文串
给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。

在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。



示例 1:

输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
示例 2:

输入:s = "a"
输入:1
示例 3:

输入:s = "bb"
输入: 2


提示:

1 <= s.length <= 2000
s 只能由小写和/或大写英文字母组成 */

// 思路：贪心算法
func longestPalindrome(s string) int {
	cnt := 0
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}

	for _, n := range sMap {
		cnt += n / 2 * 2
		if cnt%2 == 0 && n%2 == 1 {
			cnt++
		}
	}
	return cnt
}

// 打标志
func longestPalindromeEx(s string) int {
	cnt := 0
	flag := false
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}

	for _, n := range sMap {
		cnt += n / 2 * 2
		if n%2 == 1 {
			flag = true
		}
	}
	if flag {
		cnt++
	}
	return cnt
}

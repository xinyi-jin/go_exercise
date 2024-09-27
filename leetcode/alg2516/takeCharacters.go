package alg2516

/* 2516. 每种字符至少取 K 个
中等
相关标签
相关企业
提示
给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。

你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。



示例 1：

输入：s = "aabaaaacaabc", k = 2
输出：8
解释：
从 s 的左侧取三个字符，现在共取到两个字符 'a' 、一个字符 'b' 。
从 s 的右侧取五个字符，现在共取到四个字符 'a' 、两个字符 'b' 和两个字符 'c' 。
共需要 3 + 5 = 8 分钟。
可以证明需要的最少分钟数是 8 。
示例 2：

输入：s = "a", k = 1
输出：-1
解释：无法取到一个字符 'b' 或者 'c'，所以返回 -1 。


提示：

1 <= s.length <= 105
s 仅由字母 'a'、'b'、'c' 组成
0 <= k <= s.length */

// 思路：滑动窗口
func takeCharacters(s string, k int) int {
	// 校验是否存在
	tmp := [3]int{}
	for _, c := range s {
		tmp[c-'a']++
	}
	if tmp[0] < k || tmp[1] < k || tmp[2] < k {
		return -1
	}
	pre := 0
	max := 0
	for i, c := range s {
		c -= 'a'
		tmp[c]--
		for tmp[c] < k {
			tmp[s[pre]-'a']++
			pre++
		}
		max = maxNum(max, i-pre+1)
	}

	return len(s) - max
}
func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package alg3442

/* 3442. 奇偶频次间的最大差值 I
简单
相关标签
premium lock icon
相关企业
提示
给你一个由小写英文字母组成的字符串 s。

请你找出字符串中两个字符 a1 和 a2 的出现频次之间的 最大 差值 diff = freq(a1) - freq(a2)，这两个字符需要满足：

a1 在字符串中出现 奇数次 。
a2 在字符串中出现 偶数次 。
返回 最大 差值。



示例 1：

输入：s = "aaaaabbc"

输出：3

解释：

字符 'a' 出现 奇数次 ，次数为 5 ；字符 'b' 出现 偶数次 ，次数为 2 。
最大差值为 5 - 2 = 3 。
示例 2：

输入：s = "abcabcab"

输出：1

解释：

字符 'a' 出现 奇数次 ，次数为 3 ；字符 'c' 出现 偶数次 ，次数为 2 。
最大差值为 3 - 2 = 1 。


提示：

3 <= s.length <= 100
s 仅由小写英文字母组成。
s 至少由一个出现奇数次的字符和一个出现偶数次的字符组成。 */

// 思路：map计数,奇数取最大值，偶数取最小值，差值即结果
func maxDifference(s string) int {
	var js, os int
	cntMap := make(map[rune]int)
	for _, v := range s {
		cntMap[v]++
	}
	for _, v := range cntMap {
		if v%2 == 0 {
			if os == 0 {
				os = v
			} else {
				os = min(os, v)
			}
		} else {
			js = max(js, v)
		}
	}

	return js - os
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

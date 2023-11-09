package alg2609

/* 2609. 最长平衡子字符串
提示
简单
61
相关企业
给你一个仅由 0 和 1 组成的二进制字符串 s 。

如果子字符串中 所有的 0 都在 1 之前 且其中 0 的数量等于 1 的数量，则认为 s 的这个子字符串是平衡子字符串。请注意，空子字符串也视作平衡子字符串。

返回  s 中最长的平衡子字符串长度。

子字符串是字符串中的一个连续字符序列。



示例 1：

输入：s = "01000111"
输出：6
解释：最长的平衡子字符串是 "000111" ，长度为 6 。
示例 2：

输入：s = "00111"
输出：4
解释：最长的平衡子字符串是 "0011" ，长度为  4 。
示例 3：

输入：s = "111"
输出：0
解释：除了空子字符串之外不存在其他平衡子字符串，所以答案为 0 。


提示：

1 <= s.length <= 50
'0' <= s[i] <= '1' */

// 拆分思路
func findTheLongestBalancedSubstring(s string) int {
	res := 0
	str := ""
	for i, v := range s {
		str += string(v)
		if str[len(str)-1] == '1' {
			res = max(res, calcLen(str))
		} else if i == 0 || s[i-1] == '1' {
			str = "0"
		}
	}
	return res
}

// 传入参数格式，0必须在1之前
func calcLen(str string) int {
	zeroCnt, oneCnt := 0, 0
	for _, v := range str {
		if v == '0' {
			zeroCnt++
		} else {
			oneCnt++
		}
	}
	return min(zeroCnt, oneCnt) * 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 三叶解法
func findTheLongestBalancedSubstringEx(s string) int {
	n, idx, res := len(s), 0, 0
	for idx < n {
		a, b := 0, 0
		for idx < n && s[idx] == '0' {
			a++
			idx++
		}
		for idx < n && s[idx] == '1' {
			b++
			idx++
		}
		res = max(res, min(a, b)*2)
	}
	return res
}

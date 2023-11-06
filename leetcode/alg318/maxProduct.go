package alg318

/* 318. 最大单词长度乘积
中等
447
相关企业
给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。



示例 1：

输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
输出：16
解释：这两个单词为 "abcw", "xtfn"。
示例 2：

输入：words = ["a","ab","abc","d","cd","bcd","abcd"]
输出：4
解释：这两个单词为 "ab", "cd"。
示例 3：

输入：words = ["a","aa","aaa","aaaa"]
输出：0
解释：不存在这样的两个单词。


提示：

2 <= words.length <= 1000
1 <= words[i].length <= 1000
words[i] 仅包含小写字母 */

// 按位运算
func maxProduct(words []string) int {
	n := len(words)
	res := 0
	for i := 0; i < n; i++ {
		curBit := getBitValue(words[i])
		for j := i + 1; j < n; j++ {
			nexBit := getBitValue(words[j])
			if curBit&nexBit == 0 {
				res = max(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func getBitValue(str string) int {
	bitValue := 0
	for _, v := range str {
		bitValue |= 1 << (v - 'a')
	}
	return bitValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 时间优化
func maxProductEx(words []string) int {
	res := 0
	maskMap := make(map[int]int)
	for _, v := range words {
		curBit := getBitValue(v)

		// 重复字母字符串，取最长（"abc"/"aabbcc"）
		if len(v) > maskMap[curBit] {
			maskMap[curBit] = len(v)
		}
	}
	for i, m := range maskMap {
		for j, n := range maskMap {
			if i&j == 0 {
				res = max(res, m*n)
			}
		}
	}
	return res
}

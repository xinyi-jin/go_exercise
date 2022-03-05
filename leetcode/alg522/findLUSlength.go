package alg522

/* 522. 最长特殊序列 II
给定字符串列表 strs ，返回 它们中 最长的特殊序列 。如果最长特殊序列不存在，返回 -1 。

最长特殊序列 定义如下：该序列为某字符串 独有的最长子序列（即不能是其他字符串的子序列）。

 s 的 子序列可以通过删去字符串 s 中的某些字符实现。

例如，"abc" 是 "aebdc" 的子序列，因为您可以删除"aebdc"中的下划线字符来得到 "abc" 。"aebdc"的子序列还包括"aebdc"、 "aeb" 和 "" (空字符串)。


示例 1：

输入: strs = ["aba","cdc","eae"]
输出: 3
示例 2:

输入: strs = ["aaa","aaa","aa"]
输出: -1


提示:

2 <= strs.length <= 50
1 <= strs[i].length <= 10
strs[i] 只包含小写英文字母 */

// 暴力解 ，找出所有子字符串，计数为1的即是所求
func findLUSlength(strs []string) int {
	allChildStrs := make(map[string]int)
	for _, s := range strs {
		for i := 0; i < 1<<len(s); i++ {
			t := ""
			for j := 0; j < len(s); j++ {
				if ((i >> j) & 1) != 0 {
					t += string(s[j])
				}
			}
			allChildStrs[t]++
		}
	}
	res := -1
	for s, v := range allChildStrs {
		if v == 1 {
			res = max(res, len(s))
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

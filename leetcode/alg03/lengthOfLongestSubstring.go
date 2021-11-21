package alg03

/* 3. 无重复字符的最长子串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。



示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0 */

/* 解题思路：
滑动窗口：
依次滑动所有的求取所有的情况
程序流程如下所示：
str "abcedf"
a
ab
abc
abcd
abcde
abcdef
*/
func lengthOfLongestSubstring(s string) int {
	sm := map[byte]int{}
	maxLength, left, n := 0, -1, len(s)
	for i := 0; i < n; i++ {
		if i > 0 {
			delete(sm, s[i-1])
		}
		for left+1 < n && sm[s[left+1]] == 0 {
			sm[s[left+1]]++
			left++
		}
		if curLength := left - i + 1; curLength > maxLength {
			maxLength = curLength
		}
	}
	return maxLength
}

/* 解题思路：
暴力求解：查找所有子字符串，逐个判断其中元素是否重复
跑测试用例会超时 */
func lengthOfLongestSubstringBaoLi(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	ss := make([]string, 0)
	maxLength, index := 0, 0
	for i := 0; i < len(s)-1; i++ {
		for j := i; j < len(s); j++ {
			ss = append(ss, s[i:j+1])
			index++
		}
	}
	for _, s := range ss {
		if isRepet(s) && len(s) > maxLength {
			maxLength = len(s)
		}
	}
	return maxLength
}

func isRepet(s string) bool {
	sm := make(map[string]int)
	for _, v := range s {
		if _, ok := sm[string(v)]; !ok {
			sm[string(v)] = 1
		} else {
			return false
		}
	}
	return true
}

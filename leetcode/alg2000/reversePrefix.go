package alg2000

import "strings"

/* 2000. 反转单词前缀
给你一个下标从 0 开始的字符串 word 和一个字符 ch 。找出 ch 第一次出现的下标 i ，反转 word 中从下标 0 开始、直到下标 i 结束（含下标 i ）的那段字符。如果 word 中不存在字符 ch ，则无需进行任何操作。

例如，如果 word = "abcdefd" 且 ch = "d" ，那么你应该 反转 从下标 0 开始、直到下标 3 结束（含下标 3 ）。结果字符串将会是 "dcbaefd" 。
返回 结果字符串 。



示例 1：

输入：word = "abcdefd", ch = "d"
输出："dcbaefd"
解释："d" 第一次出现在下标 3 。
反转从下标 0 到下标 3（含下标 3）的这段字符，结果字符串是 "dcbaefd" 。
示例 2：

输入：word = "xyxzxe", ch = "z"
输出："zxyxxe"
解释："z" 第一次也是唯一一次出现是在下标 3 。
反转从下标 0 到下标 3（含下标 3）的这段字符，结果字符串是 "zxyxxe" 。
示例 3：

输入：word = "abcd", ch = "z"
输出："abcd"
解释："z" 不存在于 word 中。
无需执行反转操作，结果字符串是 "abcd" 。


提示：

1 <= word.length <= 250
word 由小写英文字母组成
ch 是一个小写英文字母 */

// 思路:遍历查找，找到后直接翻转前边字符串，否则直接返回
func reversePrefix(word string, ch byte) string {
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			s := []byte(word)
			for left := 0; left < i; left++ {
				s[left], s[i] = s[i], s[left]
				i--
			}
			return string(s)
		}
	}
	return word
}

func reversePrefixEx(word string, ch byte) string {
	right := strings.IndexByte(word, ch)
	if right < 0 {
		return word
	}
	s := []byte(word)
	for left := 0; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
	return string(s)
}

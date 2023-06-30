package alg2490

import (
	"strings"
)

/* 2490. 回环句
句子 是由单个空格分隔的一组单词，且不含前导或尾随空格。

例如，"Hello World"、"HELLO"、"hello world hello world" 都是符合要求的句子。
单词 仅 由大写和小写英文字母组成。且大写和小写字母会视作不同字符。

如果句子满足下述全部条件，则认为它是一个 回环句 ：

单词的最后一个字符和下一个单词的第一个字符相等。
最后一个单词的最后一个字符和第一个单词的第一个字符相等。
例如，"leetcode exercises sound delightful"、"eetcode"、"leetcode eats soul" 都是回环句。然而，"Leetcode is cool"、"happy Leetcode"、"Leetcode" 和 "I like Leetcode" 都 不 是回环句。

给你一个字符串 sentence ，请你判断它是不是一个回环句。如果是，返回 true ；否则，返回 false 。



示例 1：

输入：sentence = "leetcode exercises sound delightful"
输出：true
解释：句子中的单词是 ["leetcode", "exercises", "sound", "delightful"] 。
- leetcode 的最后一个字符和 exercises 的第一个字符相等。
- exercises 的最后一个字符和 sound 的第一个字符相等。
- sound 的最后一个字符和 delightful 的第一个字符相等。
- delightful 的最后一个字符和 leetcode 的第一个字符相等。
这个句子是回环句。
示例 2：

输入：sentence = "eetcode"
输出：true
解释：句子中的单词是 ["eetcode"] 。
- eetcode 的最后一个字符和 eetcode 的第一个字符相等。
这个句子是回环句。
示例 3：

输入：sentence = "Leetcode is cool"
输出：false
解释：句子中的单词是 ["Leetcode", "is", "cool"] 。
- Leetcode 的最后一个字符和 is 的第一个字符 不 相等。
这个句子 不 是回环句。 */

// 模拟
func isCircularSentence(sentence string) bool {
	n := len(sentence)
	if sentence[0] != sentence[n-1] {
		return false
	}
	arr := strings.Split(sentence, " ")
	for i := 0; i < len(arr)-1; i++ {
		l := len(arr[i])
		if arr[i][l-1] != arr[i+1][0] {
			return false
		}
	}
	return true
}

// 栈 因为单词长度可能为1 像“a b”这种使用栈操作会误判
func isCircularSentenceEx(sentence string) bool {
	stack := make([]byte, 0)
	arr := strings.Split(sentence, " ")
	for _, v := range arr {
		for i := 0; i < 2; i++ {
			cur := byte(' ')
			value := v[0]
			if len(stack) > 0 {
				cur = stack[len(stack)-1]
			}
			if i > 0 {
				value = v[len(v)-1]
			}
			if value != cur {
				// 压栈
				stack = append(stack, value)
			} else {
				// 弹栈
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

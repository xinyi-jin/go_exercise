package alg917

import (
	"unicode"
)

/* 917. 仅仅反转字母
给你一个字符串 s ，根据下述规则反转字符串：

所有非英文字母保留在原有位置。
所有英文字母（小写或大写）位置反转。
返回反转后的 s 。



示例 1：

输入：s = "ab-cd"
输出："dc-ba"
示例 2：

输入：s = "a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入：s = "Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"


提示

1 <= s.length <= 100
s 仅由 ASCII 值在范围 [33, 122] 的字符组成
s 不含 '\"' 或 '\\' */

func reverseOnlyLetters(s string) string {
	temp := []rune(s)
	last := len(temp) - 1
	for i := 0; i < len(temp); i++ {
		if isABC(temp[i]) {
			for j := last; j > i; j-- {
				if isABC(temp[j]) {
					temp[i], temp[j] = temp[j], temp[i]
					last = j - 1
					break
				}
			}
		}
	}
	return string(temp)
}

func isABC(r rune) bool {
	return (unicode.IsLower(r) || unicode.IsUpper(r)) && !unicode.IsSpace(r)
}

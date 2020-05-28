package main

import (
	"fmt"
	"strconv"
	"strings"
)

/* 给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func decodeString(s string) string {
	digit, indexSlice := 0, make([]int, 2)
	for i := 0; i < len(s); i++ {
		if indexSlice[0] == 0 && string(s[i]) == "[" {
			indexSlice[0] = i
		}
		if string(s[i]) == "]" {
			indexSlice[1] = i
		}
	}
	// indexSlice[0] = strings.Index(s, "[")
	// indexSlice[1] = strings.LastIndex(s, "]")
	fmt.Println("index", indexSlice[0], indexSlice[1])

	if indexSlice[0] < 1 {
		return s
	}
	digit = indexSlice[0] - 1 //数字下标
	encodeStr := s[indexSlice[0]+1 : indexSlice[1]]
	decodeStr := ""
	if digit > 0 {
		decodeStr += s[:digit]
	}

	for j := 0; j < digit; j++ {
		decodeStr += encodeStr
	}
	if indexSlice[1] != len(s) {
		decodeStr += string(s[indexSlice[1]+1:])
	}

	return decodeString(decodeStr)

}

func decodeStringEx(s string) string {
	stack := make([]string, 0)
	var preStr string
	for i := range s {

		switch s[i] {
		case '[':
			stack = append(stack, preStr)
			preStr = ""
		case ']':

			stack = extend(stack)

		default:

			if s[i] >= '0' && s[i] <= '9' {

				preStr = preStr + s[i:i+1]

			} else {

				sLen := len(stack)
				//上一个也是字母字符串，合并 merge  // "3[a2[c]d]"  "accdaccdaccd" 对应第二个]
				if sLen > 0 && (stack[sLen-1][0] < '0' || stack[sLen-1][0] > '9') {
					stack[sLen-1] = stack[sLen-1] + s[i:i+1]
				} else {
					stack = append(stack, s[i:i+1])
				}
			}

		}

	}

	var result string
	for i := range stack {
		result = result + stack[i]
	}
	return result
}

func extend(stack []string) []string {
	sLen := len(stack)
	str := stack[sLen-1]
	numStr := stack[sLen-2]
	num, _ := strconv.Atoi(numStr)
	extendStr := strings.Repeat(str, num)

	//前面也是字符串 就合并
	if sLen > 2 && (stack[sLen-3][0] < '0' || stack[sLen-3][0] > '9') {
		extendStr = stack[sLen-3] + extendStr
		stack[sLen-3] = extendStr
		stack = stack[:sLen-2]
	} else {
		stack[sLen-2] = extendStr
		stack = stack[:sLen-1]
	}
	return stack
}

func main() {
	str := "3[a1[d]]2[bc]"
	res := decodeStringEx(str)
	fmt.Println("res", res)
}

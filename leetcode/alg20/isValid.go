package main

import "fmt"

/* 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。 */

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	str := []string{}
	str = append(str[0:], string(s[0]))

	for i := 1; i < len(s); i++ {
		v := string(s[i])
		l := len(str)
		if l == 0 {
			str = append(str, v)
			continue
		}
		if isMatch(str[l-1], v) {
			str = str[:l-1]
		} else {
			str = append(str, v)
		}
	}
	if len(str) != 0 {
		return false
	}

	return true
}

func isMatch(a, b string) bool {
	if a == "(" && b == ")" || a == "{" && b == "}" || a == "[" && b == "]" {
		return true
	}
	return false
}

func main() {
	s := "[]([]){}"
	res := isValid(s)
	fmt.Println("res:", res)
}

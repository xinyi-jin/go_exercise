package main

import "fmt"

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	var (
		nullStr, result string
	)

	if 0 == len(strs) {
		return nullStr
	} else if 1 == len(strs) {
		return strs[0]
	}
	strTemp := strs[0]
	minLen := len(strs[0])

	// 记录最短字符的长度
	for _, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
		}
	}
	for i := 0; i < minLen; i++ {
		flag := false
		for j := 1; j < len(strs); j++ {
			if strTemp[i] != strs[j][i] {
				flag = true
			}
		}
		if flag {
			break
		}
		result += string(strTemp[i])
	}
	return result
}

func main() {
	// str := []string{"dog", "racecar", "car"}
	// str := []string{"flower", "flow", "flight"}
	str := []string{}
	// str := []string{"a"}
	// str := []string{"a", "a"}
	// str := []string{"aa", "ab"}
	res := longestCommonPrefix(str)
	fmt.Printf("result %s", res)
}

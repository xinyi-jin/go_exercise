package leetcode

import (
	"strconv"
)

/* 给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。 */

func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 思路 直接使用内置函数，string 转 int 进行计算后 再转换成string (数值会越界)
func AddBinary1(a string, b string) string {
	m, _ := strconv.ParseInt(a, 2, 64)
	n, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(int64(m+n), 2)
}

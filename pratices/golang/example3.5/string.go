package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello ,世界"
	fmt.Println("len", len(s))
	fmt.Println(utf8.DecodeRuneInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	var n int
	for range s {
		n++
	}
	fmt.Println("字符串的字符个数：", n)
	fmt.Println(utf8.RuneCountInString(s))

	s1 := "1-=+23.21/zhuhe/pic.gif"

	fmt.Println("hasSuffix", HasSuffix(s1, ".gif"))

	fmt.Println("hasPrefix", HasPrefix(s1, "/zhuhe"))

	fmt.Println("baseName", BaseName(s1))

	fmt.Println("comma", Comma(s1))
	fmt.Println("comma1", Comma1(s1))

	s2 := "good luck day"
	s3 := "good dayluck "
	fmt.Println("compar", Compa(s2, s3))

}

// 是否包含前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 是否包含后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 是否包含子串
func Contains(s, c string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], c) {
			return true
		}

	}
	return false
}

// 路径下取文件名称
func BaseName(s string) string {
	index := strings.LastIndex(s, "/")
	s = s[index+1:]

	index = strings.LastIndex(s, ".")
	s = s[:index]
	return string(s)
}

// 对于字符串每隔3位，添加一个 ，号
// 递归版本
func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return Comma(s[:n-3]) + "," + s[n-3:]

}

// 对于字符串每隔3位，添加一个 ，号
// 非递归版本
func Comma1(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		index := len(s) % 3
		if i == index || (i-index)%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", s[i:i+1])

	}
	return buf.String()
}

//  编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func Compa(s1 string, s2 string) bool {
	flag := false

	n := 0

	if len(s1) != len(s2) {
		return flag
	}

	for _, v := range s1 {
		if strings.Contains(s2, string(v)) {
			n++
			continue
		}
		return flag
	}

	if len(s2) == n {
		flag = true
	}

	return flag
}

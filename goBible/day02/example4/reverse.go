package main

import "fmt"

func main() {
	reverse2("hello world")
}

/* 对输入的字符串倒叙输出 */
func reverse1(str string) {
	fmt.Println("before reverse：", str)
	var result string
	strLength := len(str)
	for i := 0; i < strLength; i++ {
		result += string(str[strLength-i-1])
	}
	fmt.Println("after reverse：", result)
}

func reverse2(str string) {
	fmt.Println("before reverse：", str)

	result := make([]byte, 0)
	strLength := len(str)
	for i := 0; i < strLength; i++ {
		result = append(result, str[strLength-i-1])
	}

	fmt.Println("after reverse：", string(result))
}

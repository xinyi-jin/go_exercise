package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := []string{"t", "t", "t", "t", "t1", "t2", "t2", "t", "t", "mb"}
	bytes := []byte{'a', ' ', ' ', ' ', 's', 'b', 'c', ' ', ' ', ' ', 'p'}
	// rev := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// remove(str)
	fmt.Println("before remove:", str)
	fmt.Println("after remove:", remove(str))
	b := TrimSlice(bytes)

	for _, v := range b {
		fmt.Printf("%c,", v)
	}

	// rotate(rev)
	// fmt.Println("rev:", rev)

}

// 重写reverse函数，使用数组指针代替slice。
func reverse(arr *[6]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 编写一个rotate函数，通过一次循环完成旋转。
func rotate(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
}

//  写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func remove(str []string) []string {
	for i := 1; i < len(str)-1; i++ {
		if str[i-1] != str[i] {
			continue
		}
		str = append(str[:i-1], str[i:]...)
		i--
	}
	return str

}

// 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func TrimSlice(bytes []byte) []byte {
	for i := 1; i < len(bytes)-1; i++ {
		if unicode.IsSpace(rune(bytes[i-1])) && unicode.IsSpace(rune(bytes[i])) {
			bytes = append(bytes[:i-1], bytes[i:]...)
			i--
		}
	}
	return bytes
}

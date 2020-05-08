package main

import "fmt"

func main() {
	//testSlice2()
	//testRune()
	testArray()
}


type slice struct {
	ptr *[100]int
	len int
	cap int
}

func newSlice(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[3] = 100
	return
}

func modify1(s [5]int) {
	s[3] = 123456
	return
}

func testSlice1() {
	var s slice
	s = newSlice(s, 10)

	s.ptr[0] = 10
	modify(s)

	fmt.Println(s.ptr)
}
func testSlice2() {
	var s = [5]int{1, 2, 3, 4, 5}

	modify1(s)

	fmt.Println(s)
}

func testRune() {
	s := "hello 中国好基友"

	s1 := []rune(s)

	s1[5] = '猪'

	str := string(s1)
	fmt.Println(str)
}

func testArray() {
	s := "hell o every bady tonight !!!"

	s1 := []byte(s)

	fmt.Println(s1)
}

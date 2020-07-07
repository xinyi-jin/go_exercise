package main

import "fmt"

func main() {
	testModify()
}

func testModify() {
	str := "朱贺ello world"
	s1 := []rune(str)
	s1[0] = 'p'
	str = string(s1)

	fmt.Println(str)

	var a map[string]string
	a = make(map[string]string, 10)

	a["123"] = "zhuhe"
	fmt.Println(a["123"])

	m1 := make(map[string]map[string]int, 10)
	m1["key1"] = make(map[string]int, 10)

	m1["key1"]["name"] = 12

	fmt.Println(m1["key1"]["name"])
}

package main

import "fmt"

func main() {
	test()
}

func test() {
	s1 := new([]int)
	*s1 = make([]int, 5)
	(*s1)[0] = 10
	fmt.Println(s1)

	s2 := make([]int, 10)
	s2[0] = 100
	fmt.Println(s2)

}

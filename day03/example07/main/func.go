package main

import "fmt"

func main() {

	c := addSum

	res := opFunc(c, 100, 200)

	fmt.Println("result:", res)
}

type add_func func(int, int) int

func addSum(a int, b int) int {

	return a + b
}

func opFunc(op add_func, a int, b int) int {
	return op(a, b)
}

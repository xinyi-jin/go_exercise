package main

import "fmt"

func main() {
	var p = f()
	*p++

	fmt.Println(*p)

	fmt.Println(f() == f())
}

var v int

func f() *int {
	v = 1
	return &v
}

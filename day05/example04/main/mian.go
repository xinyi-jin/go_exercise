package main

import "fmt"

type a int

func main() {
	var b = 12
	var a a = 24

	var c int
	c = int(a)

	fmt.Println(b, a)
	fmt.Println(c)
}

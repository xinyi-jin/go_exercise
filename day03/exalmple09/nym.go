package main

import "fmt"

func main() {
	/* result := func(a int, b int) int {
		return a + b
	}(100, 200) */

	result := func(a int, b int) int {
		return a + b
	}

	fmt.Println("result：", result(100, 200))
}

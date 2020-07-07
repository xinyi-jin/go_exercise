package main

import "fmt"

func main() {
	num := calc(3)
	fmt.Println(num)
}

/* n的阶乘 */
func calc(n int) int {
	if n == 1 {
		return 1
	}
	return calc(n-1) * n
}

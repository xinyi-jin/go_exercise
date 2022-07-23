package main

import (
	"fmt"

	pop "go_exercise/pratices/golang/example2.6/popcount"
)

func main() {
	c := pop.PopCount(1234567)
	fmt.Println("二进制中1的个数:", c)
}

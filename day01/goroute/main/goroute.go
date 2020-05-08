package main

import "fmt"

func test_goroute(a int) {
	fmt.Println("a:", a)
}

/* 加法运算 */
func add(a int, b int, ch chan int) {
	ch <- (a + b)
}

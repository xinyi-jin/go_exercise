package main

import "fmt"

func main() {

	//sum := add(1, 2, 3, 4, 5, 6, 7, 8, 9)
	//fmt.Println("sum:",sum)

	jisuan()
}

func add(a int, args ...int) (sum int) {
	sum = a
	for _, v := range args {
		sum += v
	}
	return
}

func jisuan() {
	var gameScore int = 16
	var times int = 3
	for i := 1; i <=times; i++ {
		gameScore *= 2
	}
	fmt.Println(gameScore)
}

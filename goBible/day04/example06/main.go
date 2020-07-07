package main

import "fmt"

func main() {
	// var arr [5]int = [5]int{1, 2, 3, 4, 5}
	var arr [4]int = [...]int{1, 2, 3, 4}

	for _, v := range arr {
		fmt.Println(v)
	}
}

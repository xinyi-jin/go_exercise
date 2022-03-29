package main

import (
	"fmt"
	"go_exercise/leetcode/beard/reback"
)

func main() {

	testData := [][]int{
		{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
		},
	}

	for _, v := range testData {
		huxi := reback.CalcHuXi(v, 9)
		fmt.Printf("huxi %v\n", huxi)
	}
}

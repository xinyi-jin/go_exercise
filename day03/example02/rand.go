package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var input int
	flag := false
	rand := rand.Intn(100)
	fmt.Println("请输入您猜的数字：")
	for {
		fmt.Scanf("%d\n", &input)
		switch {
		case input == rand:
			fmt.Println("you are right..")
			flag = true
		case input > rand:
			fmt.Println("you are bigger...")
		case input < rand:
			fmt.Println("you are small..")
		}

		if flag {
			break
		}
	}
}

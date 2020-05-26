// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	/* 	flag := false
// 	   	if flag == false {
// 	   		fmt.Println("true")
// 		   } */
// 	num := 0
// 	start := time.Now().Unix()
// 	for {
// 		input := randNum()

// 		fmt.Printf("%d  ", input+1)

// 		if num > 5000 {
// 			fmt.Println("break")
// 			break
// 		}
// 		num++
// 		// time.Sleep(1 * time.Second)

// 	}

// 	end := time.Now().Unix()

// 	fmt.Println("total times:", end-start)
// }

// // 实现骰子功能
// func randNum() int {/*  */
// 	rand.Seed(1 * time.Now().Unix())
// 	return rand.Intn(6)
// }

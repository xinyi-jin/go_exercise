package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := 100
	b := 200

	swap(&a, &b)

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println("rand:", a)
	}

}

/* init初始化种子 */
func init() {
	rand.Seed(time.Now().Unix())
}

/* 交换两个数字的位置 */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

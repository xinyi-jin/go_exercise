package main

import "fmt"

/* 计算一个数所有的和的情况 */
func num(a int) {
	for i := 0; i < a; i++ {
		if i > a/2 {
			return
		}
		fmt.Printf("%d+%d=%d", i, a-i, a)
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/* 根据输入的数字，找出所有满足条件的数值 a方=b方+c方+d方 */
func main() {
	var str string
	for _, v := range os.Args {
		str = v
	}
	num, _ := strconv.Atoi(str)
	// var num int = 100
	for a := 2; a <= num; a++ {
		for b := 2; b < a; b++ {
			for c := 2; c < a; c++ {
				for d := 2; d < a; d++ {
					if a*a*a == b*b*b+c*c*c+d*d*d {
						fmt.Printf("%d方=%d方+%d方+%d方\n", a, b, c, d)
						log.Printf("%d方=%d方+%d方+%d方\n", a, b, c, d)
					}
				}
			}
		}
	}
}

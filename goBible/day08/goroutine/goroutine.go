package main

import (
	"fmt"
	"runtime"
	"time"
)

func Test() {
	var i int
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
}
func main() {
	// go Test()
	// for {
	// 	fmt.Println("i running in main")
	// 	time.Sleep(time.Second)
	// }

	num := runtime.NumCPU()
	fmt.Println("cpu num:", num)
	runtime.GOMAXPROCS(num)
}

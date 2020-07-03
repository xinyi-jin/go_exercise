package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 200; i++ {
		ch <- i
		fmt.Println("i", i)
	}
}

func read(ch chan int) {
	for {
		num := <-ch
		fmt.Println("num:", num)
	}
}
func main() {
	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan)
	time.Sleep(10 * time.Second)
}

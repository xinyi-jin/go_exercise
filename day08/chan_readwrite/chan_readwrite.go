package main

import "fmt"

func read(ch <-chan int, exitChan chan struct{}) {
	var num int
	for i := 0; i < 100; i++ {
		num = <-ch
		fmt.Println("read num:", num)
	}
	var a struct{}
	exitChan <- a
}
func write(ch chan<- int, exitChan chan struct{}) {
	for u := 0; u < 100; u++ {
		ch <- u
	}
	var a struct{}
	exitChan <- a
}

func main() {
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go write(ch, exitChan)
	go read(ch, exitChan)

	for i := 0; i < 2; i++ {
		<-exitChan
	}

}

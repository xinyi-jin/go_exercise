package main

import (
	"fmt"
	"sync"
)

func read(ch <-chan int, wg *sync.WaitGroup) {
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("read num:", num)
	}
}
func write(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for u := 0; u < 100; u++ {
		ch <- u
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 10)
	wg.Add(2)

	go write(ch, &wg)
	go read(ch, &wg)

	wg.Done()
	fmt.Println("wait start")
	wg.Wait()
	fmt.Println("wait close")

	wgFunc()
}

package main

import "fmt"

func main() {
	/* for i := 0; i < 100; i++ {
		go test_goroute(i)
	}
	time.Sleep(10 * time.Second) */

	ch := make(chan int, 1)
	go add(10, 5, ch)
	sum := <-ch
	fmt.Println("10+5=", sum)
}

package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	for {
		var rb int
		rb, ok := <-ch
		if !ok {
			fmt.Println("channel is close")
			return
		}
		fmt.Println(rb)
	}
}

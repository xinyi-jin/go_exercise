package main

import "fmt"

func main() {
	for {
		var rb int
		rb, ok := <-ch
		if ok == false {
			fmt.Println("channel is close")
			return
		}
		fmt.Println(rb)
	}
}

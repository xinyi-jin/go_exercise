package main

import (
	"fmt"
	"sync"
	"time"
)

func wgFunc() {
	sayHello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("%v goroutine start ...\n", id)
		time.Sleep(2)
		fmt.Printf("%v goroutine exit ...\n", id)
	}

	var wg sync.WaitGroup
	const N = 5
	wg.Add(N)
	for i := 0; i < N; i++ {
		go sayHello(&wg, i)
	}

	fmt.Println("waiting for all goroutine ")
	wg.Wait()
	fmt.Println("All goroutines finished!")
}

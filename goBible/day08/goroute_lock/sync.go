package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]int)
	lock sync.Mutex
)

type task struct {
	n int
}

// 计算阶乘
func calc(t *task) {
	var sum int
	sum = 1
	for i := 1; i <= t.n; i++ {
		sum *= i
	}
	fmt.Println(t.n, sum)
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}
func main() {

	for i := 2; i < 10; i++ {
		t := &task{n: i}
		go calc(t)
	}
	time.Sleep(1 * time.Second)

	/* lock.Lock()
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	lock.Unlock() */
}

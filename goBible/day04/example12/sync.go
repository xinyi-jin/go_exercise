package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// var lock sync.Mutex

var rLock sync.RWMutex
var count int64

func main() {
	s := make([]int, 10)
	s[0] = 10

	for i := 0; i < 20000; i++ {
		go modify(s)
		atomic.AddInt64(&count, 1)
	}

	for i := 0; i < 100; i++ {
		// time.Sleep(time.Second)
		rLock.RLock()
		fmt.Println(s[0])
		rLock.RUnlock()
	}

	fmt.Println(atomic.LoadInt64(&count))

}

func init() {
	rand.Seed(time.Now().Unix())
}

func modify(s []int) {
	rLock.Lock()
	s[0] = rand.Intn(100)
	rLock.Unlock()
}

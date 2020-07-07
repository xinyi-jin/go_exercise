package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()

	fmt.Println("耗时：", end-start)
}

func test() {
	time.Sleep(time.Microsecond * 100)
}

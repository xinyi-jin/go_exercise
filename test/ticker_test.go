package test

import (
	"fmt"
	"testing"
	"time"
)

// 因为本身返回的tick值就在一个管道中，阻塞当前线程，所以还未到外层for循环阻塞线程
func Test_TickerInFor(t *testing.T) {
	i, j := 0, 0
	for {
		tick := time.Tick(time.Second * 5)
		for c := range tick {
			fmt.Printf("%v %v %v\n", c, i, j)
			j++
		}
		time.Sleep(time.Second * 3)
		i++
	}
}

// 使用管道接收tick
func Test_TickerInForEx(t *testing.T) {
	i, j := 0, 0
	for {
		// go func() {
		select {
		case tick := <-time.Tick(time.Second * 1):
			fmt.Printf("%v %v %v\n", tick, i, j)
			j++
		}
		// }()
		time.Sleep(time.Second * 10)
		i++
	}
}

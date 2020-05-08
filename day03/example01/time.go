package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		now := time.Now()
		fmt.Println("now:", now.Format("2006-01-02 15:04:05"))

		time.Sleep(1 * time.Second)

		if now.Second() == 0 {
			break
		}

	}
}

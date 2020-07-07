package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("index :%v,value:%v \n", i, v)
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "c", "", "please input your path")

	flag.Parse()
	fmt.Println("string path:", confPath)
}

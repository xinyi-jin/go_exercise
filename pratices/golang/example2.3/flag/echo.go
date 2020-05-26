package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "this is bool")
var s = flag.String("s", " ", "tab space")

func main() {

	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}

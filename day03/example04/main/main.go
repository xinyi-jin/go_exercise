package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world abc"
	result := strings.Replace(str,"hello","nihao",1)

	fmt.Println(result)

	count := strings.Count(str,"abc")

	fmt.Println(count)

	strs := strings.Split(str," ")

	for _, v := range strs {
		fmt.Println(v)
	}

}

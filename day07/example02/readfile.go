package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	读取文件的几种方式：
	ioutile.ReadAll
	bufio.NewReader().ReadString()
*/
func main() {
	file, err := os.Open("h:/test.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read file err:", err)
	}
	fmt.Println("read file :", str)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("G:/test.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line []byte
	for {
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line = append(line, data...)
		if !prefix {
			fmt.Printf("data:%s\n", string(line))
			line = line[:]
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("h:/test.txt", os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	// gfile, err := gzip.NewReader(file)
	// if err != nil {
	// 	fmt.Println("open gzip err", err)
	// 	return
	// }
	// // defer gfile.Close()

	// reader := bufio.NewReader(gfile)
	// str, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("read gfile err:", err)
	// 	return
	// }

	writer := bufio.NewWriter(file)
	writer.WriteString("\n")
	num, err := writer.WriteString("lalalalallllllalsdfklj")
	if err != nil {
		fmt.Println("writer file err:", err)
	}
	fmt.Println("writer success num :", num)
	writer.Flush()

	// fmt.Println("read file :", str)

	CopyFile("H:/test.txt", "H:/main.txt")
}

func CopyFile(openFile, writeFile string) (written int64, err error) {

	rd, err := os.Open(openFile)
	if err != nil {
		fmt.Println("read file err:", err)
		return 0, err
	}
	defer rd.Close()

	wr, err := os.OpenFile(writeFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("write file err:", err)
		return 0, err
	}
	defer wr.Close()

	return io.Copy(wr, rd)

}

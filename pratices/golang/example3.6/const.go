package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math"
	"os"
)

const PI = 3.1415926

const (
	_ = 1 << (10 * iota)
	a
	b
	c
	d
)

func main() {
	fmt.Printf("a %b\n", a)
	fmt.Println("a", a)
	fmt.Printf("b %b\n", b)
	fmt.Printf("c %b\n", c)
	fmt.Printf("d %b\n", d)
	fmt.Println("pi", PI)
	fmt.Printf("Math.PI:%0.48f\n", math.Pi)
	fmt.Println("______________________")

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

}

// 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
func Print() []byte {
	var c1 []byte
	str := os.Args[0:1]
	switch str {
	case nil:
		c1 = sha256.Sum256([]byte("x"))
	case SHA384:
		c1 = sha384.Sum384([]byte("x"))
	case SHA512:
		c1 = sha512.Sum512([]byte("x"))
	}
	return c1
}

/* 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
	设计思路：PopCount函数计算的是一个值中含有某个bit的数目，for循环遍历这两个sha256其中含有bit为0的就是不同的bit位 */
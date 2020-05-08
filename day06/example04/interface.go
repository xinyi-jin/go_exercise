package main

import "fmt"

func main() {

	// TypeAssert("shdf")

	TypesWord(12)
}

// 类型断言
func TypeAssert(a interface{}) {
	// c := a.(int)
	c, ok := a.(int)
	if !ok {
		fmt.Println("type failed")
	}
	fmt.Println("c", c)
}

func TypesWord(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case int:
			fmt.Println("int", v)
		case string:
			fmt.Println("string", v)
		case bool, float64:
			fmt.Println("bool int", v)
		}

	}
}

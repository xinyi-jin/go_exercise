package main

import (
	"fmt"
	"sort"
)

func main() {
	//testMap()
	//testMap1()
	testMapSort()
}

func testMap() {
	var m map[string]int

	m = make(map[string]int)
	m["123"] = 123

	fmt.Println(m)
	delete(m, "123")
	fmt.Println("key：123 对应的value", m["123"])

}

func testMap1() {
	var m map[string]map[int]int = make(map[string]map[int]int)

	m["001"] = make(map[int]int)
	m["001"][123] = 123
	//m["002"][456] = 458
	//m["003"][789] = 12
	//m["004"][147] = 1
	//m["005"][258] = 123

	for k, v := range m {
		fmt.Println(k, v)
		for k1, v1 := range m["001"] {
			fmt.Println(k1, v1)
		}
	}

}

func testMapSort() {
	var (
		m map[int]int
		mslice []int
	)

	if m == nil {
		m = make(map[int]int)
	}

	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4

	//map直接遍历输出的是无序的
	//for key, value := range m {
	//	fmt.Println("key:",key,"value:",value)
	//}

	for key, _ := range m {
		mslice = append(mslice, key)
	}

	sort.Ints(mslice)

	for _, value := range mslice {
		fmt.Println("sort after ,key:",value,"value:",m[value])
	}
}

func testMapReserve()  {
	var (
		m1 map[int]string
		m2 map[string]int
	)

	if m1 == nil {
		m1=make(map[int]string)
	}
	if m2 == nil {
		m2=make(map[string]int)
	}


}

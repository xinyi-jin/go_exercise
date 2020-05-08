package main

import (
	"fmt"
	"sort"
)

func main() {
	//testSortInt()
	//testSortString()
	testSortFloat()
}

func testSortInt() {
	var ints = []int{1, 2, 3, 4, 598, 78}

	sort.Ints(ints)

	fmt.Println(ints)
}

func testSortString() {
	var ints = []string{"f", "g", "a"}

	sort.Strings(ints)

	fmt.Println(ints)
}

func testSortFloat() {
	var ints = []float64{0.25, 2.59, 0.01, 2.53}

	sort.Float64s(ints)

	fmt.Println(ints)
}

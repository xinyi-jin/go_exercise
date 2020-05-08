package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	var stus SortByStuAge
	for i := 0; i < 10; i++ {
		stu := Student{
			Id:   rand.Intn(20),
			Name: fmt.Sprintf("stu%d", rand.Intn(1000)),
			Age:  rand.Intn((100)),
		}

		stus = append(stus, stu)
	}

	for _, v := range stus {
		fmt.Println("sort before:", v)
	}
	// 排序
	sort.Sort(stus)

	fmt.Println("\n")

	for _, v := range stus {
		fmt.Println("sort after:", v)
	}
}

type Student struct {
	Id   int
	Name string
	Age  int
}

/* 实现Sort接口，用来排序 */
type SortByStuAge []Student

func (a SortByStuAge) Len() int           { return len(a) }
func (a SortByStuAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByStuAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

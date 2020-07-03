package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)

	stu := student{name: "stu01"}

	stuChan <- &stu

	var stu01 interface{}
	stu01 = <-stuChan

	var stu02 *student
	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("can not convert")
		return
	}
	fmt.Println(stu02)

	var mapChan chan map[string]string

	mapChan = make(chan map[string]string, 10)

	m := make(map[string]string, 10)
	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%ds", i)
		m[s] = s + "00"
	}

	var mm map[string]string
	mm = make(map[string]string, 5)
	mapChan <- m
	mm = <-mapChan

	fmt.Println("mm", mm)
}

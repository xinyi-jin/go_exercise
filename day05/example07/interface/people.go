package main

import "fmt"

func main() {
	var stu1 *Student = &Student{
		stuNo:    1,
		stuName:  "lili",
		stuAge:   20,
		stuClass: "大一",
	}

	var p PeoPle

	p = stu1
	p.run()
	p.sleep(2)

}

type PeoPle interface {
	run()
	sleep(t int)
}

type Student struct {
	stuNo    int
	stuName  string
	stuAge   int
	stuClass string
}

func (s Student) run() {
	fmt.Println(s.stuName, "run")
}

func (s Student) sleep(t int) {
	fmt.Println(s.stuName, "sleep total:", t, "hours")
}

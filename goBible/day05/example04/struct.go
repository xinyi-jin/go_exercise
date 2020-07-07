package main

import "fmt"

func main() {
	// var s Student
	// s.stuId = 1
	// s.stuSex = 1
	// s.stuName = "朱贺"
	// s.stuAge = 22
	// s.stuClass = "大四"
	s := Student{
		stuAge:  22,
		stuId:   1,
		stuName: "zhuhe",
	}

	stu1 := Student{
		stuAge:  22,
		stuId:   1,
		stuName: "xiaobenzhu",
	}
	s.next = &stu1
	stu2 := Student{
		stuAge:  22,
		stuId:   1,
		stuName: "benzhu ",
		// next:nil
	}
	stu1.next = &stu2

	/* 	var p *Student = &s
	   	for p != nil {
	   		fmt.Println(p.stuName)
	   		p = p.next
	   	} */

	conver(&s)

	// fmt.Println(&s.stuAge)
	// fmt.Println(&s.stuId)

}

/* 学生信息 */
type Student struct {
	stuId    int
	stuSex   int
	stuName  string
	stuAge   int
	stuClass string
	next     *Student
}

func conver(p *Student) {
	for p != nil {
		fmt.Println(p.stuName)
		p = p.next
	}
}

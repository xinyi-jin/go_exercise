package main

import "fmt"

type Student struct {
	name string
	age  int
	next *Student
}

func main() {
	var stu *Student = new(Student)
	stu.name = "heiil"
	stu.age = 23
	// insertHead(&stu)
	insertTail(stu)

	delNode(stu)
	cover(stu)

}

func cover(p *Student) {
	for p != nil {
		fmt.Println(p.name)
		p = p.next
	}
}

/* 头部插入 */
func insertHead(s **Student) {
	var head = s
	for i := 0; i < 2; i++ {
		stu := &Student{
			name: "xiaoming",
			age:  18,
		}
		stu.next = *head
		*head = stu
	}
}

/* 尾部插入 */
func insertTail(s *Student) {
	var tail = s
	for i := 0; i < 2; i++ {
		stu := &Student{
			name: "xiaoming",
			age:  18,
		}
		tail.next = stu
		tail = stu
	}
}

/* 删除一个节点 */
func delNode(s *Student) {
	var prev *Student = new(Student)
	for s != nil {
		if s.name == "heiil" {
			prev.next = s.next
		}
		prev = s.next
		prev = s
	}
}

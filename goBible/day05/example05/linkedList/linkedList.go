package main

import "fmt"

/* 单链表的增删改查 */
func main() {
	//定义初始链表头
	var head *Student = new(Student)
	head.name = "hello"
	//头部插入
	// insertHead(&head)

	//尾部插入
	insertTail(head)

	delNode(head)
	trans(head)
}

/* 学生类型 */
type Student struct {
	name  string
	age   int
	score float32
	next  *Student
}

/* 遍历输出 */
func trans(p *Student) {
	for p != nil {
		fmt.Println(p)
		p = p.next
	}
}

/* 头部插入 */
func insertHead(head **Student) {
	for i := 0; i < 2; i++ {
		stu := &Student{
			name:  "xiaoming",
			age:   18,
			score: 112.5,
		}
		stu.next = *head
		*head = stu
	}
}

/* 尾部插入 */
func insertTail(tail *Student) {
	for i := 0; i < 10; i++ {
		stu := &Student{
			name:  "liuxiang",
			age:   i + 20,
			score: 121.02,
		}

		tail.next = stu
		tail = stu
	}
}

func delNode(p *Student) {
	var prev *Student = p

	for p != nil {
		if p.age == 22 {
			prev.next = p.next
		}
		//保存上一个节点
		prev = p
		//不满足条件切换到下一个节点
		p = p.next
	}
}

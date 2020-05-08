package main

import "fmt"

func main() {
	var stu = new(Student)
	stu.init("zhuhe", 23)

	stuInfo := stu.getInfo()
	fmt.Println("student info:", *stu)

	fmt.Println("getInfo Student:", stuInfo)

	fmt.Println(stu)

}

type Student struct {
	Name string
	Age  int
}

type integer int

func (i integer) setValues() {

}

// 此处传递的是指针，会直接修改指针地址指向的值
func (s *Student) init(name string, age int) {
	s.Name = name
	s.Age = age
}

// 此处s 指的是s的一份拷贝，修改属性内容，并不会修改s原本的属性值
func (s Student) getInfo() Student {
	s.Name = "小笨猪"
	return s
}

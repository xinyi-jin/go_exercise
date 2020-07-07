package main

import "fmt"

func main() {
	var root *Student = new(Student)

	root.Name = "string"
	root.Age = 20
	root.Score = 0.02

	var left1 = &Student{
		Name:  "int",
		Age:   18,
		Score: 0.01,
	}
	root.left = left1

	var right1 = &Student{
		Name:  "float32",
		Age:   23,
		Score: 0.09,
	}

	root.right = right1

	traverse(root)

}

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func traverse(r *Student) {
	if r == nil {
		return
	}
	fmt.Println(r.Name)

	traverse(r.left)
	traverse(r.right)
}

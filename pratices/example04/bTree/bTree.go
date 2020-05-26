package main

import "fmt"

func main() {

	var s = &Student{
		Name: "name",
		Age:  21,
	}
	var s2 = &Student{
		Name: "name1",
		Age:  21,
	}
	s.left = s2
	var s3 = &Student{
		Name: "name2",
		Age:  21,
	}
	s.right = s3

	converse(s)

}

type Student struct {
	Name  string
	Age   int
	left  *Student
	right *Student
}

func converse(s *Student) {
	if s == nil {
		return
	}
	fmt.Printf("\t%s\n", s.Name)
	converse(s.left)
	converse(s.right)
}

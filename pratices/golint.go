package main

import (
	"fmt"
	"go_work/pratices/golint/moudle"
)

func main() {
	student := moudle.NewStudent(1, "bnen")

	fmt.Println("student name:", student.getName())
}

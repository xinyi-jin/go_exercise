package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s op=%s createTime=%s message=%s", p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open("test.xls")
	defer file.Close()
	if err != nil {
		return &PathError{
			path:       "test.slx",
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	return nil
}

func main() {
	err := Open("f:sf/sad/f.xls")
	if v, ok := err.(*PathError); ok {
		fmt.Println("get path error,", v)
		return
	}
}

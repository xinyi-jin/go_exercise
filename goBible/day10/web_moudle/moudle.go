package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	age  int
}

func main() {
	t, err := template.ParseFiles("H:/go/src/goProject/day10/web_moudle/index.html")
	if err != nil {
		fmt.Println("parse file err,", err)
	}

	p := Person{Name: "zhuhe", age: 22}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("err,", err)
	}
}

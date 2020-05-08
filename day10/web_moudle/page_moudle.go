package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var myTemplate *template.Template

type Person struct {
	Name string
	age  int
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err,", err)
		return
	}
	return
}

func userinfo(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "hello ,world!")
	p := Person{Name: "zhuhe", age: 22}
	if err := myTemplate.Execute(w, p); err != nil {
		fmt.Println("err,", err)
	}
}

func main() {
	initTemplate("H:/go/src/goProject/day10/web_moudle/index.html")
	http.HandleFunc("/user/info", userinfo)
	http.ListenAndServe("localhost:8000", nil)
}

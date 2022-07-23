package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var myTemplate *template.Template

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

func parseFiles() {
	t, err := template.ParseFiles("H:/go/src/go_exercise/goBible/day10/web_moudle/index.html")
	if err != nil {
		fmt.Println("parse file err,", err)
	}

	p := Person{Name: "zhuhe", age: 22}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("err,", err)
	}
}

func main() {
	parseFiles()
	initTemplate("H:/go/src/go_exercise/goBible/day10/web_moudle/index.html")
	http.HandleFunc("/user/info", userinfo)
	http.ListenAndServe("localhost:8000", nil)
}

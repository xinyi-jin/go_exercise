package main

import (
	"log"
	"net/http"
	"text/template"
)

type Person struct {
	Name    string
	Age     int
	Address string
	Sex     string
}

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf(" Create template failed :err%x", err)
		return
	}
	// 渲染模板
	tmpl.Execute(w, "小笨猪")
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf("Create template failed :err%x", err)
		return
	}
	// 渲染模板
	Person := &Person{
		Name:    "小笨猪",
		Age:     22,
		Address: "ZHUMADIAN",
		Sex:     "woman",
	}
	tmpl.Execute(w, *Person)
}

func f3(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	tmpl, err := template.ParseFiles("./map.tmpl")
	if err != nil {
		log.Printf("Create template failed :err %x", err)
		return
	}
	// 渲染模板
	var hobby = make(map[string]string)
	hobby["h1"] = "movie"
	hobby["h2"] = "running"
	hobby["h3"] = "talking"

	tmpl.Execute(w, hobby)
}

func f4(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	tmpl, err := template.ParseFiles("./map.tmpl")
	if err != nil {
		log.Printf("Create template failed :err %x", err)
		return
	}
	Person := &Person{
		Name:    "小笨猪",
		Age:     22,
		Address: "ZHUMADIAN",
		Sex:     "woman",
	}
	var hobby = make(map[string]string)
	hobby["h1"] = "movie"
	hobby["h2"] = "running"
	hobby["h3"] = "talking"
	var mapArr = make(map[string]interface{})
	mapArr["hobby"] = hobby
	mapArr["person"] = *Person

	// 渲染模板
	tmpl.Execute(w, mapArr)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/f2", f2)
	http.HandleFunc("/f3", f3)
	http.HandleFunc("/f4", f4)
	log.Println("web server start. ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("web server start failed: err%x", err)
	}
	log.Println("web server start sucess. ")
}

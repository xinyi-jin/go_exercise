package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

// Person 人类
type Person struct {
	Name    string
	Age     int
	Address string
	Sex     string
}

// 解析数据：字符串
func f1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf(" Create template failed :%v", err)
		return
	}
	// 渲染模板
	tmpl.Execute(w, "小笨猪")
}

// 解析数据 struct
func f2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Printf("Create template failed :%v", err)
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

// 解析数据 map
func f3(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	tmpl, err := template.ParseFiles("./map.tmpl")
	if err != nil {
		log.Printf("Create template failed :v", err)
		return
	}
	// 渲染模板
	var hobby = make(map[string]string)
	hobby["h1"] = "movie"
	hobby["h2"] = "running"
	hobby["h3"] = "talking"

	tmpl.Execute(w, hobby)
}

// 解析数据 map 接口类型
func f4(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	tmpl, err := template.ParseFiles("./map.tmpl")
	if err != nil {
		log.Printf("Create template failed :%v", err)
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

// 自定义模板函数
func f5(w http.ResponseWriter, r *http.Request) {
	/* // 解析模板
	tmpl, err := template.New("mytmpl.tmpl").ParseFiles("./mytmpl.tmpl")
	if err != nil {
		log.Printf("Create template failed :%v", err)
		return
	}
	// 渲染模板
	tmpl.Execute(w, "xiaozhuzhu") */

	htmlByte, err := ioutil.ReadFile("./mytmpl.tmpl")
	if err != nil {
		log.Println("read html failed, err:", err)
		return
	}
	var praise = func(s string) (res string, err error) {
		return s + "高富帅呦", nil
	}
	// 解析模板
	t, err := template.New("mytmpl.tmpl").Funcs(template.FuncMap{"praise": praise}).Parse(string(htmlByte))
	if err != nil {
		log.Printf("parse template failed :%v", err)
		return
	}

	// 渲染模板
	t.Execute(w, "臭猪猪")
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/f2", f2)
	http.HandleFunc("/f3", f3)
	http.HandleFunc("/f4", f4)
	http.HandleFunc("/f5", f5)
	log.Println("web server start. ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("web server start failed: err%v", err)
	}
	log.Println("web server start sucess. ")
}

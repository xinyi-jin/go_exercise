package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var stu = &Student{
		id:      1,
		Name:    "zhuhe",
		age:     23,
		address: "驻马店泌阳",
	}

	jsonBytes, _ := json.Marshal(*stu)

	fmt.Println(string(jsonBytes))

}

type Student struct {
	id      int    `json:"id"`
	Name    string `json:"name"`
	age     int    `json:"age"`
	address string `json:"homeAddress"`
}

/* type Student struct {
	id      int
	name    string
	age     int
	address string
} */

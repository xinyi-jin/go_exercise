package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var stu = &Student{
		Id:      1,
		Name:    "zhuhe",
		Age:     23,
		Address: "驻马店泌阳",
	}

	jsonBytes, _ := json.Marshal(*stu)

	fmt.Println(string(jsonBytes))

}

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"homeAddress"`
}

/* type Student struct {
	id      int
	name    string
	age     int
	address string
} */

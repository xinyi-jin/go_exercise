package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func main() {
	// TestStruct()
	// TestInt()
	// TestMap()
	TestSlice()

}

func TestStruct() {
	user1 := &User{
		UserName: "user1",
		NickName: "zhuhe",
		Age:      22,
		Birthday: "1997",
		Sex:      "女",
		Email:    "iezhuhe@163.com",
		Phone:    "15226048650",
	}

	bytes, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json marshal err:", err)
	}

	bs, _ := json.MarshalIndent(user1, "", "\t")
	fmt.Println("json:", string(bytes))
	fmt.Println(string(bs))
}
func TestInt() {
	var a int = 10
	i, err := json.Marshal(a)
	if err != nil {
		fmt.Println("json marshal err", err)
	}
	fmt.Println("int json :", string(i))
}

func TestMap() {
	var m map[string]string
	m = make(map[string]string)

	m["name"] = "zhuhe"
	m["age"] = "22"
	i, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal err", err)
	}
	fmt.Println("int json :", string(i))
}

func TestSlice() {
	var s []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "zhuhe"
	m["age"] = 22
	s = append(s, m)

	m["name"] = "jinxinyi"
	m["age"] = "22"
	s = append(s, m)

	i, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal err", err)
	}
	fmt.Println("slice json :", string(i))

}

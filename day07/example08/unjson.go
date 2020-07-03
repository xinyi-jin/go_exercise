package main

import (
	"encoding/json"
	"log"
)

type User struct {
	UserName string
	Age      int
	Email    string
	Phone    string
}

var user *User

func jsonMarshal() []byte {
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("json marshal err, ", err)
	}

	return data
}

func jsonUnMarshal(data []byte) User {
	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		log.Println("json unmarshal err, ", err)
	}
	return u
}

func init() {
	user = &User{
		UserName: "zhuhe",
		Age:      22,
		Email:    "iezhuhe@163.com",
		Phone:    "15226048650",
	}
}

func main() {
	data := jsonMarshal()
	log.Println("user byte ", string(data))
	u := jsonUnMarshal(data)
	log.Println("user ", u)
}

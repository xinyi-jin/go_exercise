package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Student ...
type Student struct {
	ID       uint8
	Name     string
	Age      uint8
	Address  string
	homeWork *HomeWork
}

// HomeWork ...
type HomeWork struct {
	Chinese string
	English string
}

func (g *Student) methodName() {
	g.Name = "1234564687"
}

// GetHomeWork ...
func (s *Student) GetHomeWork() *HomeWork {
	return s.homeWork
}

/*
	go使用gob序列化与反序列化
	ps:struct中的fields首字母大写，不然无法访问
*/
func main() {
	s1 := &Student{1, "zhuhe", 22, "ZhengZhouHeNan", &HomeWork{}}
	//序列化
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer) //创建编码器
	err1 := encoder.Encode(s1)         //编码
	if err1 != nil {
		log.Println("err1", err1)
	}
	fmt.Printf("序列化后：%x\n", buffer.Bytes())

	byteDec := buffer.Bytes()
	var s2 Student
	// 反序列化
	decoder := gob.NewDecoder(bytes.NewReader(byteDec))

	err2 := decoder.Decode(&s2)
	if err2 != nil {
		log.Println("err2", err2)
	}
	fmt.Println("反序列化后：", s2)

	// golint test
	var name = "zhuhe"
	fmt.Println("name:", name)
}

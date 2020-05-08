package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	stu := &Student{
		id:      1,
		name:    "zhuhe",
		age:     19,
		address: "zhengzhou",
	}

	res, _ := json.Marshal(stu)
	fmt.Println("json data:", res)

	Test(*stu)
}

func Test(in interface{}) {
	/* s := reflect.ValueOf(in)
	fmt.Println("value:", s)

	t := s.Type()
	fmt.Println(t)

	v := s.String()
	fmt.Println(v) */
	// 获取value对象
	val := reflect.ValueOf(in)

	// 获得value对象的类型	ps：Type()和Kind()获得对象的区别点：Type()获得是具体类型 main.Student	Kind()获得的是对象类型 struct
	ty := val.Type()

	// 获得value对象的类型
	tys := val.Kind()
	if tys != reflect.Struct {
		fmt.Println("tys is not struct")
	}

	fmt.Println("type :", ty)
	fmt.Println("tys :", tys)
	fmt.Println("reflect.Struct :", reflect.Struct)
	fmt.Println(val.String())

	// 字段数量
	fmt.Println("struct field number :", val.NumField())
	// 方法数量		ps:获取方法数量的时候，需要设置方法的名称大写，才能统计到方法数量
	fmt.Println("struct methord number :", val.NumMethod())
	var ref []reflect.Value
	val.Method(0).Call(ref)

}

type Student struct {
	id      int    `json:id`
	name    string `json:name`
	age     int
	address string
}

func (s Student) Set(id int, name string, age int, address string) {
	s.id = id
	s.name = name
	s.age = age
	s.address = address
}
func (s Student) Print() {
	fmt.Println("test student Print")
}

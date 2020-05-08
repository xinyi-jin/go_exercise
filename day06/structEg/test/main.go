package main

import (
	"fmt"
	"goProject/day06/structEg/moudle"
)

func main() {
	stu := moudle.NewStudent(1, "zhuhe", 22, "zhengzhou")
	fmt.Println("student", *stu)
	fmt.Printf("id:%d,name:%s,age:%d,address:%s", stu.GetId(), stu.GetName(), stu.GetAge(), stu.GetAddress())

}

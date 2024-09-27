package main

import (
	"fmt"
	"strconv"
	"time"
)

type Data struct {
	Name string
}

func main() {
	fmt.Println("Hello World!!!")

	var f = [5]func(){}
	for i := 0; i < 5; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}
	for _, v := range f {
		v()
	}

	data := []*Data{
		{Name: "haha"},
		{Name: "wawa"},
		{Name: "nana"},
		{Name: "gugu"},
	}
	data = append(data[:0], data[2:]...)
	for _, v := range data {
		println(v.Name)
	}
	data = append(data, &Data{Name: "biubiuqiu"})
	for _, v := range data {
		println("new ", v.Name)
	}

	nowHourMinute := time.Unix(time.Now().Unix(), 0).Format("1504")
	nowInt, err := strconv.ParseInt(nowHourMinute, 10, 64)
	if err != nil {
		print("err:", err)
	}
	print(nowHourMinute, nowInt)
}

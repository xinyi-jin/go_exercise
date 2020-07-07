package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("get res err:", err)
		return
	}

	resdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read res data err:", err)
		return
	}
	fmt.Println("res data:", string(resdata))
}

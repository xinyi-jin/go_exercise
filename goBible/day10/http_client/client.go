package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func head() {
	for _, v := range url {
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}
		resp, err := c.Head(v)

		// resp, err := http.Head(v)
		if err != nil {
			fmt.Println("http head err:", err)
			continue
		}
		fmt.Println("respone status:", resp.Status)
	}
}

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

	head()
}

package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {

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

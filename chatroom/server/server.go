package main

import (
	"fmt"
	"net"
)

// 广播消息管道
var message = make(chan string)

func start(network, addr string) {
	listener, err := net.Listen(network, addr)
	if err != nil {
		fmt.Printf("start Listen err:%v\n", err)
	}
	defer listener.Close()
	fmt.Println("server start sucess.")

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Printf("start Accept err:%v\n", err)
			}
			go func() {
				for {
					buf := make([]byte, 4096)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("start Read err:%v\n", err)
						conn.Close()
						break
					}
					if n != 0 {
						s := string(buf[:n])
						fmt.Print(s)
					}
				}
			}()
		}
	}()
	select {}
}

func main() {
	start("tcp", "127.0.0.1:80")
}

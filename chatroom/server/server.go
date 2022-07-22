package main

import (
	"fmt"
	"net"
)

// 在线用户
var OnlineMap map[string]net.Conn

func start(network, addr string) {
	OnlineMap = make(map[string]net.Conn)
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
			OnlineMap[conn.RemoteAddr().String()] = conn
			go func() {
				for {
					buf := make([]byte, 4096)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("start Read err:%v\n", err)
						delete(OnlineMap, conn.RemoteAddr().String())
						conn.Close()
						break
					}
					if n != 0 {
						s := string(buf[:n])
						fmt.Print(s)
						boardcastMsg(buf[:n])
					}
				}
			}()
		}
	}()
	select {}
}

func boardcastMsg(buf []byte) {
	// fmt.Println("boardcastMsg ", len(OnlineMap))
	for _, c := range OnlineMap {
		c.Write(buf)
	}
}

func main() {
	start("tcp", "127.0.0.1:80")
}

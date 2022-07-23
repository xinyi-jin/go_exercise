package main

import (
	"fmt"
	"go_exercise/chatroom/log"
	"net"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.NewLogger("client_log", log.DEBUG)
}

func start() {
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		logger.Errorf("net.Dial err:", err)
		return
	}
	defer conn.Close()
	logger.Infoln("connecting sucess.")

	// 客户端发送消息到服务器
	go func() {
		for {
			buffer := make([]byte, 4096)
			// 这里使用Stdin标准输入，因为scanf无法识别空格
			n, err := os.Stdin.Read(buffer)
			if err != nil {
				logger.Errorf("os.Stdin.Read err:", err)
				continue
			}
			_, _ = conn.Write(buffer[:n])
		}
	}()
	// 接收服务器发来的数据
	go func() {
		for {
			buffer := make([]byte, 4096)
			n, err := conn.Read(buffer)
			if n == 0 {
				logger.Infoln("connection is shuted.")
				return
			}
			if err != nil {
				logger.Infoln("conn.Read err:", err)
				return
			}
			fmt.Print(string(buffer[:n]))
		}
	}()
	select {}
}

func main() {
	start()
}

package main

import (
	"fmt"
	usermodule "go_exercise/chatroom/user_module"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("connecting sucess.")

	user := usermodule.NewUserInfo()
	_, _ = conn.Write([]byte(user.GetUserName()))

	/* buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println(string(buf[:n])) */

	// 客户端发送消息到服务器
	go func() {
		for {
			buffer := make([]byte, 4096)
			// 这里使用Stdin标准输入，因为scanf无法识别空格
			n, err := os.Stdin.Read(buffer)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
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
				fmt.Println("connection is shuted.")
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:", err)
				return
			}
			fmt.Print(string(buffer[:n]))
		}
	}()
	select {}
}

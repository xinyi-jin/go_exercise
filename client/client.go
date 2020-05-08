package main

import (
	"fmt"
	"net"
)

/*多人在线聊天室--客户端*/
func main() {
	//连接服务器消息
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	//错误处理
	errors(err)
	//	延迟关闭
	defer conn.Close()

	//写入传输给客户端的程序
	conn.Write([]byte("I am zhuhe, I love icecream"))
	fmt.Println("send message to server success...")

	fmt.Println("waitting for server message...")
	var arr = make([]byte, 1024)
	conn.Read(arr)

	fmt.Printf("server receive message ...%s\n", string(arr))

}

/*处理错误*/
func errors(err error) {
	if err != nil {
		panic(err)
	}
}

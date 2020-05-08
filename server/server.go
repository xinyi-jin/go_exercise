package main

import (
	"fmt"
	"net"
)

/*多人在线聊天室--服务器*/
func main() {
	//开启监听
	server_scoket, err := net.Listen("tcp", "127.0.0.1:80")
	//错误处理
	errors(err)
	//延迟关闭连接
	defer server_scoket.Close()

	//记录log，用于表示服务器状态
	fmt.Println("server is waitting for client message...")
	for true {
		//接收数据
		receMsg, err := server_scoket.Accept()
		errors(err)

		//启用并发，接收消息
		go prossInfo(receMsg)
	}

}

/*错误处理*/
func errors(err error) {
	if err != nil {
		panic(err)
	}
}

func prossInfo(conn net.Conn) {
	//定义字节数组用于存放数据
	var buf = make([]byte, 1024)
	//延迟关闭conn
	defer conn.Close()

	for true {
		//用于读取客户端传送的字节数组
		numOfBytes, err := conn.Read(buf)

		//如果发生错误就返回
		if err != nil {
			break
		}

		//把读取到的数据打印到控制台
		if numOfBytes != 0 {
			fmt.Printf("receive message:%s\n", string(buf))
			break
		}

	}
	fmt.Printf("input you receive message:\n", )

	var str string
	fmt.Scan(&str)

	conn.Write([]byte(str))

	fmt.Println("receive client message success..")

}

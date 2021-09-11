package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	usermodule "go_exercise/chatroom/user_module"
)

/*多人在线聊天室--客户端*/
func main() {

	//记录log，用于表示服务器状态
	fmt.Println("connect server sucessful !")
	username := usermodule.GenUserName()
	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 建立连接
		conn := connect()
		//	延迟关闭
		defer conn.Close()

		// 读取命令行输入
		str, err := inputReader.ReadString('\n')
		errors(err)

		// 发送给server
		str = username + ":" + str
		conn.Write([]byte(str))

		// 读取房间内消息
		var arr = make([]byte, 1024)
		conn.Read(arr)

		fmt.Printf("%s\n", string(arr))
	}

}

/*处理错误*/
func errors(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() net.Conn {
	//连接服务器消息
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	//错误处理
	errors(err)
	return conn
}

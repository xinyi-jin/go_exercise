package main

import (
	"fmt"
	"go_exercise/chatroom/log"
	usermodule "go_exercise/chatroom/user_module"
	"net"
)

var logger *log.Logger

type ChatRoomServer struct {
	method   string       // 连接方式 tcp websocket
	ip       string       // ip地址
	port     int32        // 端口
	listener net.Listener // 监听接口
}

func (crs *ChatRoomServer) Start() {
	var err error
	addr := fmt.Sprintf("%s:%d", crs.ip, crs.port)
	crs.listener, err = net.Listen(crs.method, addr)
	if err != nil {
		logger.Errorf("(crs *ChatRoomServer) Start() Listen err:%v\n", err)
	}
	defer crs.listener.Close()
	logger.Infof("server start sucess.\n")
	go crs.Receiver()
	select {}
}

func (crs *ChatRoomServer) Receiver() {
	for {
		conn, err := crs.listener.Accept()
		if err != nil {
			logger.Errorf("(crs *ChatRoomServer) Receiver(%s) Accept err:%v\n", conn.RemoteAddr().String(), err)
		}

		SocketManagerSingleton.AddSocekt(conn)
		usermodule.UserManagerSingleton.AddUser("", conn)
		user := usermodule.UserManagerSingleton.GetUser(conn)
		msg := fmt.Sprintf("welcome to chatroom %s\n", user.GetUserName())
		otherMsg := fmt.Sprintf("%s enter chatroom\n", user.GetUserName())
		SocketManagerSingleton.SendToClient(conn, msg)
		SocketManagerSingleton.BoardcastMsgOthers(conn, otherMsg)
		go crs.ProcessMsg(conn)
	}
}

func (crs *ChatRoomServer) ProcessMsg(conn net.Conn) {
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			logger.Errorf("(crs *ChatRoomServer) ProcessMsg Read(%s) err:%v\n", conn.RemoteAddr().String(), err)
			crs.Exit(conn)
			break
		}
		if n != 0 {
			preMsg := fmt.Sprintf("(crs *ChatRoomServer) ProcessMsg(%s)", conn.RemoteAddr().String())
			user := usermodule.UserManagerSingleton.GetUser(conn)
			switch {
			case n > 8 && string(buf[:8]) == "/rename ":
				newName := ""
				if string(buf[n-2:n]) == "\r\n" {
					newName = string(buf[8 : n-2])
				} else if string(buf[n-1:]) == "\n" {
					newName = string(buf[8 : n-1])
				} else {
					newName = string(buf[8:n])
				}
				user.ModifyUserName(newName)
				SocketManagerSingleton.SendToClient(conn, "rename sucess")
			case string(buf[:5]) == "/exit":
				crs.Exit(conn)
			default:
				msg := fmt.Sprintf("%s:%s", user.GetUserName(), string(buf[:n]))
				logger.Info(preMsg + msg)
				SocketManagerSingleton.BoardcastMsgOthers(conn, msg)
			}
		}
	}
}

func (crs *ChatRoomServer) Exit(conn net.Conn) {
	logger.Infoln("server close connection")
	conn.Close()
	SocketManagerSingleton.DelSocekt(conn)
	usermodule.UserManagerSingleton.DelUser(conn)
}

func init() {
	logger = log.NewLogger("server_log", log.DEBUG)
}

func NewChatRoomServer(method, ip string, port int32) *ChatRoomServer {
	return &ChatRoomServer{
		method: method,
		ip:     ip,
		port:   port,
	}
}

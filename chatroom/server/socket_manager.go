package main

import (
	"net"
)

var SocketManagerSingleton = &SocketManager{
	onlineMap: make(map[string]*Socket),
}

type SocketManager struct {
	onlineMap map[string]*Socket // 在线用户
}

func (sm *SocketManager) GetOnlineMap() map[string]*Socket {
	return sm.onlineMap
}

func (sm *SocketManager) BoardcastMsg(msg string) {
	for _, socket := range sm.onlineMap {
		socket.WriteMsg(msg)
	}
}

func (sm *SocketManager) BoardcastMsgOthers(conn net.Conn, msg string) {
	key := conn.RemoteAddr().String()
	for k, socket := range sm.onlineMap {
		if k == key {
			continue
		}
		socket.WriteMsg(msg)
	}
}

func (sm *SocketManager) SendToClient(conn net.Conn, msg string) {
	socket := NewSocket(conn)
	socket.WriteMsg(msg)
}

func (sm *SocketManager) AddSocekt(conn net.Conn) {
	socket := NewSocket(conn)
	key := conn.RemoteAddr().String()
	sm.onlineMap[key] = socket
	logger.Debugf("(sm *SocketManager) AddSocekt(%s)\n", conn.RemoteAddr().String())
}

func (sm *SocketManager) DelSocekt(conn net.Conn) {
	key := conn.RemoteAddr().String()
	delete(sm.onlineMap, key)
	logger.Debugf("(sm *SocketManager) DelSocekt(%s)\n", key)
}

package usermodule

import (
	"net"
)

var UserManagerSingleton = &UserManager{
	onlineUser: make(map[string]*User),
}

type UserManager struct {
	onlineUser map[string]*User
}

func (um *UserManager) AddUser(name string, conn net.Conn) {
	user := NewUser(name)
	key := conn.RemoteAddr().String()
	um.onlineUser[key] = user
	logger.Debugf("(um *UserManager) AddUser(%s) %v\n", conn.RemoteAddr().String(), *user)
}

func (um *UserManager) GetUser(conn net.Conn) *User {
	key := conn.RemoteAddr().String()
	return um.onlineUser[key]
}

func (um *UserManager) DelUser(conn net.Conn) {
	logger.Debugf("(um *UserManager) DelUser(%s) %v\n", conn.RemoteAddr().String(), um.onlineUser)
	key := conn.RemoteAddr().String()
	delete(um.onlineUser, key)
}

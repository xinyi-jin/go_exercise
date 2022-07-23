package main

import (
	"net"
)

type Socket struct {
	conn net.Conn // 连接
}

func (s *Socket) GetConn() net.Conn {
	return s.conn
}

func (s *Socket) WriteMsg(msg string) {
	s.conn.Write([]byte(msg))
	logger.Debugf("(s *Socket) WriteMsg(%s):%s", s.conn.RemoteAddr().String(), msg)
}

func NewSocket(conn net.Conn) *Socket {
	if conn == nil {
		return nil
	}
	return &Socket{
		conn: conn,
	}
}

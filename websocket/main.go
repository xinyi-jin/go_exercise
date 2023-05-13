package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"log"

	"github.com/gorilla/websocket"
)

type ProtoHeader struct {
	Len     uint16 //包长度
	Seq     uint16 //包序号
	LogicNo uint32 //逻辑号
}

type PacketHeader struct {
	EncodeType int16
	PacketId   int16
}

type RWBuffer struct {
	pheader ProtoHeader
	seq     uint16
	buf     []byte
}

func main() {
	var err error
	// 连接WebSocket服务器
	// wss需有下行配置
	dialer := websocket.Dialer{TLSClientConfig: &tls.Config{RootCAs: nil, InsecureSkipVerify: true}}
	conn, _, err := dialer.Dial("wss://192.168.1.106:33002/", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sndbuf := &RWBuffer{
		pheader: ProtoHeader{
			Len:     12,
			Seq:     1,
			LogicNo: 1,
		},
	}
	ioBuf := bytes.NewBuffer([]byte{})
	err = binary.Write(ioBuf, binary.LittleEndian, sndbuf.pheader.Len)
	err = binary.Write(ioBuf, binary.LittleEndian, sndbuf.pheader.Seq)
	err = binary.Write(ioBuf, binary.LittleEndian, sndbuf.pheader.LogicNo)
	if err != nil {
		return
	}
	ph := PacketHeader{
		EncodeType: int16(2),
		PacketId:   int16(-2003),
	}
	binary.Write(ioBuf, binary.LittleEndian, &ph)
	binary.Write(ioBuf, binary.LittleEndian, []byte{})
	// 发送消息
	err = conn.WriteMessage(websocket.BinaryMessage, ioBuf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// 读取消息
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Received message:", string(p), messageType)
}

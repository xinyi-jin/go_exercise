package main

func main() {
	s := NewChatRoomServer("tcp", "127.0.0.1", 80)
	s.Start()
}

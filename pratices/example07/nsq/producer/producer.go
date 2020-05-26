package main

import (
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	for i := 0; i < 10; i++ {
		sendMessage()
	}
	time.Sleep(time.Second * 10)
}

func sendMessage() {
	url := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(url, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	err = producer.Publish("testing", []byte("hello world"))
	if err != nil {
		panic(err)
	}
	producer.Stop()
}

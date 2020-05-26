package main

import (
	"fmt"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	testNSQ()
}

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func testNSQ() {
	url := "127.0.0.1:4150"

	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("testing", "struggle", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}

			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(url)
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}
	}()

	waiter.Wait()
}

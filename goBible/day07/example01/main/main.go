package main

import (
	"errors"
	"fmt"
	"goProject/day07/example01/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	// var balancer balance.Balancer
	// 测试使用：随机产生16个ip地址
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("172.16.10.%d", rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	// balancer := &balance.RandomBalance{}
	// balancer := &balance.RoundRobinBalance{}

	// 通过命令传参调用
	if len(os.Args) < 2 {
		fmt.Printf("error:%s", errors.New("os.Args error!"))
		return
	}
	balanceName := os.Args[1]

	/* switch balanceName {
	case "random":
		balancer = &balance.RandomBalance{}
	case "roundrobin":
		balancer = &balance.RoundRobinBalance{}
	default:
		balancer = &balance.RandomBalance{}
	} */

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}

		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand踩坑：伪随机数，每次运行程序的时候，第一次的随机结果都是一样的
	var num int
	fmt.Println("请输入您的数字：")
	fmt.Scanf("%d", &num)

	switch num {
	case randMath():
		fmt.Println("恭喜您，猜对了")
	default:
		fmt.Println("不好意思，您猜错了")
	}

}

func randMath() int {
	// 随机种子，按时间生成随机数
	rand.Seed(time.Now().Unix())
	math := rand.Intn(100)

	fmt.Println("rand math :", math)
	return math
}

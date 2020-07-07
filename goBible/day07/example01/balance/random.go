package balance

import (
	"errors"
	"fmt"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

func (r *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {

	if len(insts) == 0 {
		err = errors.New("no instan1ce")
		return
	}

	fmt.Println("randbalance start")

	// 随机算法
	lens := len(insts)
	index := rand.Intn(lens)

	return insts[index], nil
}

package balance

import (
	"errors"
	"fmt"
)

type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

func (r *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {

	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}

	lens := len(insts)
	/* if r.curIndex >= lens {
		r.curIndex = 0
	} */

	fmt.Println("roundrobin start")
	// 轮询算法
	inst = insts[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens

	return
}

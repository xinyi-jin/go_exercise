package balance

import (
	"fmt"
)

type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func (b *BalanceMgr) registerBalancer(name string, bar Balancer) {
	b.allBalancer[name] = bar
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not Found %s Balancer", name)
		return
	}

	inst, err = balancer.DoBalance(insts)
	return
}

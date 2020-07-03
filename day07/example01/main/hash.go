package main

import (
	"fmt"
	"goProject/day07/example01/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (h *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("no backend instance")
		return
	}

	fmt.Println("hash start")
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]

	return
}

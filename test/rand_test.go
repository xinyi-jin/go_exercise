package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Student struct {
	rander *rand.Rand
}

func Test_Scene1(t *testing.T) {
	stu := Student{}
	stu.rander = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("time now :", time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		if i > 5 {
			buildScene()
			return
		}
		buildScene2()
	}
}

func buildScene() bool {
	return buildScene2()
}
func buildScene2() bool {
	return true
}

package test

import (
	"fmt"
	_ "go_exercise/log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	readBytes, err := os.ReadFile("info_log.log")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("readBytes:%s\n", readBytes)
	}

}

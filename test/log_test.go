package test

import (
	"fmt"
	_ "go_exercise/log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	file, _ := os.Open("info_log.log")
	readBytes := make([]byte, 1024)
	_, err := file.Read(readBytes)
	// readBytes, err := os.ReadFile("info_log.log") // ReadFile 官方已废弃
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("readBytes:%s\n", readBytes)
	}

}

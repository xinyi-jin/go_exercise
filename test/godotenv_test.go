package test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// 测试godotenv
func TestRead(t *testing.T) {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("godotenv.Read err:%v", err)
	}
	log.Print(envMap)
	log.Print(os.Getenv("TEST_GODOTENV"))
	log.Print(os.Getenv("windir"))
}
func TestLoad(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("godotenv.Read err:%v", err)
	}
	log.Print(os.Getenv("TEST_GODOTENV"))
	log.Print(os.Getenv("windir"))
}

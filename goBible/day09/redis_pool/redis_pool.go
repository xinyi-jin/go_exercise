package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	c := pool.Get()
	defer c.Close()

	_, err := c.Do("Set", "zhuhe", "xiaobenzhu")
	if err != nil {
		fmt.Println(err)
		return
	}

	// r, err := redis.Int(c.Do("Get", "zhuhe"))
	r, err := redis.String(c.Do("Get", "zhuhe"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("r", r)
	pool.Close()
}

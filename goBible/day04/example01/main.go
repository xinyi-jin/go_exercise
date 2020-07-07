package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	/* file, _ := http.Get("http://www.baidu.com")

	out, _ := ioutil.ReadAll(file.Body)

	fmt.Println(string(out)) */

	/* var a []int
	a = append(a, 10, 11, 12)
	a = append(a, a...)
	fmt.Println(a) */

	for {
		err := initConfig()
		if err != nil {
			panic(err)
		}

		// test()
		time.Sleep(time.Second)

	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	a := 0
	b := 3 / a
	fmt.Println(b)
}

func initConfig() (err error) {
	return errors.New("error zero not use good")
}

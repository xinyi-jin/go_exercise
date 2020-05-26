package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Println("total time:", time.Since(start).Seconds())

}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	status := resp.Status
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// f, err := ioutil.ReadAll(resp.Body)
	// reader := strings.NewReader(string(f))
	// nbytes, err := io.Copy(os.Stdout, resp.Body) //结果输出到控制台
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) //结果不输出
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	resp.Body.Close()

	// 执行用时
	totalTime := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%2f %7d %s  %s", totalTime, nbytes, url, status)

}

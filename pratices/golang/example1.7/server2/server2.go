package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "%d count:\n", count)
	mu.Unlock()
}

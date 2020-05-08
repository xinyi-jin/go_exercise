package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}
func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
}

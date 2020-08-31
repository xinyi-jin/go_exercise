package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gin ")
}

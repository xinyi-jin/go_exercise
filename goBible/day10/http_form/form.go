package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html>
		<body>
			<form action="#" method="post" name="bar">
				<input type="text" name="in"/> 
				<br />
				<input type="text" name="in"/> 
				<input type="submit" name="Submit"/> 
			</form>
		</body>
	</html>`

func SinmpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello ,world!")
	panic("panic")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}
}

func LogPanic(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic %v", r.RemoteAddr, x)
			}
		}()
		handle(w, r)
	}
}
func main() {
	http.HandleFunc("/test1", LogPanic(SinmpleServer))
	http.HandleFunc("/test2", LogPanic(FormServer))

	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
	}
}

package main

import "fmt"

type Reader interface {
	read(src string) string
}

type Writer interface {
	write(src string, str string) bool
}

type RWriter interface {
	Reader
	Writer
}

type file struct {
	FileName string
	Str      string
}

func (f *file) read(src string) string {
	s := "read file "
	return s
}
func (f *file) write(src, str string) bool {
	fmt.Println("write file ")
	return true
}
func main() {
	// file := &file{
	// 	FileName: "testFile",
	// 	Str:      "read and writer",
	// }
	// fmt.Println(file.read("/test.text"))
	// file.write("/test.text", "t")

	var (
		f  file
		rw RWriter
	)
	rw = &f
	fmt.Println(rw.read("/test.text"))
	rw.write("/test.text", "t")
}

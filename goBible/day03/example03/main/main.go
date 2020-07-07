package main

import (
	"fmt"
	"strings"
)

func main() {

	var(
		url string
		path string
	)
	fmt.Println("请输入网址：")
	fmt.Scanf("%s",&url)
	fmt.Println("请输入路径:")
	fmt.Scanf("%s",&path)

	url = urlProcess(url)
	fmt.Println("url :",url)

	path = pathProcess(path)
	fmt.Println("path :",path)

}

func urlProcess(url string) string {
	result := strings.HasPrefix(url,"http://")

	if !result {
		url = fmt.Sprintf("http://%s",url)
	}
	return url
}

func pathProcess(path string)string  {
	result := strings.HasSuffix(path,"/")

	if !result {
		path = fmt.Sprintf("%s/",path)
	}
	return path
}

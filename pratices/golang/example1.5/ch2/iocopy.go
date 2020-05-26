package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/* 获取http地址的信息 */
func main() {
	url := "mobsec-dianhua.baidu.com/dianhua_api/open/location?tel=15226048650"
	prefix := "http://"
	if strings.HasPrefix(url, prefix) == false {
		url = prefix + url
	}
	// f, _ := http.Get("http://mobsec-dianhua.baidu.com/dianhua_api/open/location?tel=15226048650")
	// f, _ := http.Get("http://www.baidu.com")
	f, _ := http.Get(url)
	str, _ := ioutil.ReadAll(f.Body)
	byts, _ := io.Copy(os.Stdout, strings.NewReader(string(str)))
	fmt.Println(byts)

	// resp.Status 查看请求的http的状态码
	fmt.Println(f.Status)
}

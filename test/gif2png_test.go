package test

import (
	"fmt"
	"image/gif"
	"image/png"
	"net/http"
	"os"
	"testing"
)

func TestGif2Png(t *testing.T) {
	// 发起HTTP GET请求
	resp, err := http.Get("https://test-static.xiaoxinxin.com/admin/2025/11/27/01d3ea9f-f255-4ce0-92ed-dc85c1d5b9c4.gif")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 确保连接关闭

	// 读取响应体内容
	// content, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(content))

	// // 打开GIF文件
	// gifFile, err := os.Open("https://test-static.xiaoxinxin.com/admin/2025/11/27/01d3ea9f-f255-4ce0-92ed-dc85c1d5b9c4.gif")
	// if err != nil {
	// 	panic(err)
	// }
	// defer gifFile.Close()

	// 解码GIF文件
	gifImage, err := gif.DecodeAll(resp.Body)
	if err != nil {
		panic(err)
	}

	for i, v := range gifImage.Image {
		// 创建PNG文件
		pngFile, err := os.Create(fmt.Sprintf("output_%d.png", i))
		if err != nil {
			panic(err)
		}
		defer pngFile.Close()

		// 编码并保存PNG文件
		err = png.Encode(pngFile, v)
		if err != nil {
			panic(err)
		}
	}
}

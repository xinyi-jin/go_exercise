package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin
	r := gin.Default()
	// 加载模板
	r.LoadHTMLFiles("./index.html", "./index1.html")

	// 监听 /index
	r.GET("./index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 监听 /index
	r.GET("./index1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 单个文件上传
	r.POST("./upload", func(c *gin.Context) {
		// 获取上传文件
		f, err := c.FormFile("uploadfile")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 文件保存地址
			dst := "./data/" + f.Filename
			if err := c.SaveUploadedFile(f, dst); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "OK",
				})
			}
		}
	})
	// 多个文件上传
	r.POST("/uploadfiles", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		files := form.File["files"]

		for i, f := range files {
			dst := fmt.Sprintf("./data/%s_%d", f.Filename, i)
			// 逐个上传文件
			if err := c.SaveUploadedFile(f, dst); err != nil {
				log.Println("upload file error,", f.Filename, err)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	// 启动gin服务
	r.Run(":9090")
}

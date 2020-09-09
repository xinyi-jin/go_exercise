package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 使用index 会报404
	r.GET("/baidu", func(c *gin.Context) {
		// 请求重定向，浏览器url会改变
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 路由重定向
	r.GET("/router", func(c *gin.Context) {
		// 类似请求转发,浏览器url不会改变
		c.Request.URL.Path = "/hello"
		r.HandleContext(c)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello ya",
		})
	})

	r.Run(":9090")
}

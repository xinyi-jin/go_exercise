package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载路由
	router := gin.Default()
	// 加载模板
	// "./sources/index.html"
	router.LoadHTMLFiles("./sources/login.html", "./sources/index.html")
	// 监听请求， 渲染模板
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// 监听表单请求
	router.POST("/login", func(c *gin.Context) {
		// 获取表单数据  方法一  直接获取key对应的值
		// name := c.PostForm("name")
		// age := c.PostForm("age")

		// 方法二  key值不存在时，使用第二参数用作默认值
		// name := c.DefaultPostForm("names", "xinyi-jin")
		// age := c.DefaultPostForm("age", "22")

		// 方法三 返回是否存在key值的标识
		name, ok := c.GetPostForm("names")
		if !ok {
			name = "defaultName"
		}
		age, ok := c.GetPostForm("age")
		if !ok {
			age = "defaultAge"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name": name,
			"Age":  age,
		})
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 获取请求URI 参数
	router.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 启动路由服务
	router.Run(":9090")
}

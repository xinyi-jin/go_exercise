package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `form:"username" json:"username"`
	Age  string `form:"age" json:"age"`
}

// ShouldBind会按照下面的顺序解析请求中的数据完成绑定：

// 如果是 GET 请求，只使用 Form 绑定引擎（query）。
// 如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
func main() {
	// gin 初始化
	r := gin.Default()

	// 绑定参数
	// 使用GET请求时， ShouldBind只能使用form 绑定  使用postman时 要传递form参数
	r.GET("/form", func(c *gin.Context) {
		var u user
		// 绑定json参数
		if err := c.ShouldBind(&u); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": u.Name,
				"age":      u.Age,
			})
		} else {
			// 400参数信息错误
			c.JSON(http.StatusBadRequest, err.Error())
		}
	})

	// 使用POST请求时， ShouldBind 可以是JSON 也可以是 FORM, 根据前端实际传递的数据绑定
	r.POST("/form", func(c *gin.Context) {
		var u user
		// 绑定json参数
		if err := c.ShouldBind(&u); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"username": u.Name,
				"age":      u.Age,
			})
		} else {
			// 400参数信息错误
			c.JSON(http.StatusBadRequest, err.Error())
		}
	})

	// gin 启动
	r.Run(":9090")
}

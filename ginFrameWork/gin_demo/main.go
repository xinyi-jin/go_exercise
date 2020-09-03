package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	// tag 指定传给前端的json格式
	Name    string `json:"name"`
	Address string `json:"地址"`
	Hobby   string
}

func main() {
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Get",
		})
	})

	router.POST("/somePost", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "POST",
		})
	})

	router.PUT("/somePut", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "PULL",
		})
	})

	router.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "DELETE",
		})
	})

	// 返回map结构化json
	router.GET("/ginmap", func(c *gin.Context) {
		// gin.H 实现原理

		// 方法一
		// var resMap = make(map[string]interface{})
		// resMap["name"] = "xiaozhuzhu"
		// resMap["gender"] = "woman"

		// 方法二
		// resMap := map[string]interface{}{
		// 	"name":    "zhuzhu",
		// 	"address": "zhumadian",
		// 	"age":     22,
		// }

		// 方法三
		resMap := gin.H{
			"name":    "zhuzhu",
			"address": "zhumadian",
			"age":     22,
			"work":    "护士小姐姐",
		}

		c.JSON(http.StatusOK, resMap)
	})

	// 返回结构体 json
	router.GET("./ginstruct", func(c *gin.Context) {
		u := &user{
			"zhuzhu",
			"zhumadian",
			"冰激凌 炒酸奶 火锅 烤肉",
		}
		c.JSON(http.StatusOK, u)
	})

	// Get请求参数 querystring
	// 使用key - value 形式传递, 多个参数使用 & 连接
	router.GET("/user", func(c *gin.Context) {
		// n := c.Query("name")
		n, ok := c.GetQuery("name") //获取key不存在的时候，返回第二参数 false
		if !ok {                    //key不存在
			n = "defaultName"
		}
		// a := c.Query("age")
		a := c.DefaultQuery("ages", "23") //获取key不存在的时候，返回key 默认值

		c.JSON(http.StatusOK, gin.H{
			"name":          n,
			"age":           a,
			"name is exits": ok,
		})
	})

	router.Run(":9090")
}

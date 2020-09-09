package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	c := gin.Default()
	c.GET("/gg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"methord": "ok"})
	})

	// 路由组 应用场景：模块功能划分， API接口设计
	// {}可以省略，写上逻辑更清晰
	// 路由组支持嵌套语法
	group := c.Group("/user")
	{
		group.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET methord",
			})
		})

		group.GET("/hobby", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET hobby methord",
			})
		})
		group.POST("/info", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "POST methord",
			})
		})
		group.PUT("/info", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"message": "PUT methord",
			})
		})
		group.DELETE("/info", func(c *gin.Context) {
			c.JSON(http.StatusNoContent, gin.H{
				"message": "DELETE methord",
			})
		})
	}

	// 404
	c.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"404": "pages not found",
		})
	})

	c.Run(":9090")
}

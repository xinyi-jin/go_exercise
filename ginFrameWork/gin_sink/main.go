package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// gin中间件中使用goroutine
// 当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。

func main() {
	// Default源码 ，默认使用Logger 和 Recovery 中间件
	// engine := New()
	// engine.Use(Logger(), Recovery())
	r := gin.Default()

	// 获取不使用中间件的路由引擎
	// r := gin.New()
	// 定义中间件 程序花费时间
	var costTimeMW = func(c *gin.Context) {
		startTime := time.Now()
		// 执行后续的处理函数
		c.Next()
		// 不执行后续处理函数，直接返回
		// c.Abort()
		costTime := time.Since(startTime)
		fmt.Printf("cost time %v\n", costTime)
	}

	// 全局路由使用中间件
	r.Use(costTimeMW)

	// 单个路由使用中间件
	r.GET("/sink", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "sucess",
		})
	})

	// 路由组使用中间件
	// group := r.Group("/user", costTimeMW)
	// 或者使用use方法
	group := r.Group("/user")
	// group.Use(costTimeMW)
	{
		group.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"gin group": "user/info",
			})
		})
	}

	r.Run(":9090")
}

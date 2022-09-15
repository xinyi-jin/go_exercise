---
title: Gin web开发入门
date: 2020-08-29 16:45:00
tags:
- Go
- Gin
categories:
- 学习【资料整理记录】
---

- Day01 简单部署Gin框架
- 
ps：首次执行的时候我是mac默认的go 1.12版本，然后就很多错误。后来换成go 1.14.5版本，执行下方教程就好啦。

1. 安装Go语言开发环境
使用mac命令行，brew工具直接安装，代码如下：

```brew install go```

运行完成没有报错就是安装成功

2. 安装Gin框架
使用官方提供的命令安装。（前提：需要设置go全局的path变量，自行查找方案，不做过多介绍）

```go get -u github.com/gin-gonic/gin```

这里直接运行命令的话会出现 连接超时报错。因为get的资源中包含一些其他的资源引用，这部分正好处于google包中或其他国内不能访问的包中。
解决方法：
设置GOPROXY变量，用于获取这部分资源，然后get到本地。

使用 go env 命令查看 GOPROXY变量的值。
执行以下命令，打开modules 然后设置goproxy代理

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

完成上述操作后，再执行第一步的go get命令，没有报错就证明下载成功。可以到gopath目录下查找gin包。

3. 部署Gin 简单web页面
首先，构建go源码文件，导入gin包

```import "github.com/gin-gonic/gin"```
然后根据官网给的example，如下：

```
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

最后，执行go run xxx.go执行
注意：最后执行的时候大概率会报错

```
main.go:3:8: cannot find module providing package github.com/gin-gonic/gin: working directory is not part of a module
```

具体原因是因为go mod 本身的资源管理问题，在开启modules的时候，资源路径改变，导致不能正确访问到gin包。导致go代码中 import 的时候没有引入gin包。

解决方法：在当前执行代码的目录下，执行以下命令。

```
go mod init gin 
go mod edit -require github.com/gin-gonic/gin@latest
```

这个好像是类似于nodejs的npm包管理工具，了解不多不做过多描述，以免误导各位。

小技巧：使用go build 的时候可以使用-o参数指定生成的文件名称
gin_test 自定义的生成文件名称
gin_demo.go程序源代码

```go build -o gin_test gin_demo.go```

关于Gin API方面，官网提供Using GET, POST, PUT, PATCH, DELETE and OPTIONS

```
func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
	
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
```

值得一提的是，通过网址访问http://127.0.0.1:8080/ 默认使用的get请求，用其他的都会404，应该是前端请求时候要指定是post或者delete或其他的请求才能正确访问到。

好像是根据RESTful API接口风格来定义的，感兴趣的可以看下阮一峰老师的博文：
RESTful API 设计指南

在这里的话如果我们想要看其他请求的效果，可以使用postman api测试，选择请求方式，填写url地址就可以发起请求，然后看到返回的数据。


Gin框架官方文档：https://gin-gonic.com/docs/.
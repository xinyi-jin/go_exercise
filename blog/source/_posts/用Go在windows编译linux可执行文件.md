---
title: 用Go在windows编译linux可执行文件
date: 2019-12-01 16:45:00
tags:
- Go
- Windows
- Linux
categories:
- 学习【资料整理记录】
---

用Go在windows编译linux可执行文件

一、配置好windows下的Go环境
	```GOPATH   GOROOT```
	
二、进入CMD命令窗口
```
SET GOARCH=amd64
SET	GOOS=linux
```
注意： “=”两边不要留空格，否则会修改参数失败，编译出来的还是window下的可执行文件。具体原因我也不清楚，不加空格就对了。

三、编译文件
	```go build XXX.go```
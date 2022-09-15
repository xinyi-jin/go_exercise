---
title: go func
date: 2020-03-01 12:00:00
tags:
- Go
categories:
- 学习【资料整理记录】
---

## 2020.03.01

golang 匿名函数

​	可赋值给变量，使用变量名调用
```

	f:=func(){

		...

	}

	f()
```


​	闭包
```
	func(a int){

		...

	}(1)
```


​	匿名函数用作回调函数

 

​	// 遍历切片的每个元素, 通过给定函数进行元素访问
```
	func visit(list []int, f func(int)) {

  	for _, v := range list {

    	f(v)

  	}

	}
```
  // 使用匿名函数打印切片内容
```
  visit([]int{1, 2, 3, 4}, func(v int) {

    fmt.Println(v)

  })
```


​	封装匿名函数，实现动态调用
```
	var fMap = map[string]{

		"temp" : func(a int){

			...

		}

	}

 

	f := fMap["temp"]

	f(1)
```
---
title: GoLang中if临时变量的生存周期
date: 2020-03-25 16:45:00
tags:
- Go
categories:
- 学习【资料整理记录】
---

进行了下方代码进行验证

```
package main

import "fmt"

func main(){
	if n:=1+1;n!=2{
		fmt.Print("if n:",n)
	}else if n<2{
		fmt.Print("else if n:",n)
	}else {
		fmt.Print("else n:",n)
	}
	// fmt.Print("n:",n)
}

```

输出结果：
```else n:2```

总结，由代码执行结果可得：if 判断条件中临时定义的变量，其生命周期存在于整个if else语句逻辑中，不管是if分支逻辑还是else分支逻辑。注意：在if else外则直接会提示 undefined: n

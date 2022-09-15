---
title: go int类型转string类型
date: 2020-03-25 16:45:00
tags:
- Go
categories:
- 学习【资料整理记录】
---

版权声明：本文为CSDN博主「duzhenxun」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/duzhenxun/article/details/95739946

2个值做对比时要先注意他们的类型.如果一个string与一个in32不能直接对比.先要转换类型

如果使用Itoa需要一个int,使用FormatInt需要一个int64,

最简单的方法是使用fmt.Sprint(int32),但效率比较低

1. fmt.Sprint(i) 比较慢

```
func Sprint(a ...interface{}) string {
    p := newPrinter()
    p.doPrint(a)
    s := string(p.buf)
    p.free()
    return s
}
```

2. strconv.Itoa(int(i))

```
func Itoa(i int) string {
    return FormatInt(int64(i), 10)
}
```

3. 


```
strconv.FormatInt(int64(i), 10)
func FormatInt(i int64, base int) string {
    _, s := formatBits(nil, uint64(i), base, i < 0, false)
    return s
}
```

4. string转成int：

```int, err := strconv.Atoi(string)```

5. string转成int64：

```int64, err := strconv.ParseInt(string, 10, 64)```
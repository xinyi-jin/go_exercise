---
title: 去除windows脚本cmd小黑框
date: 2020-03-17 16:45:00
tags:
- Windows
categories:
- 学习【资料整理记录】
---

去除windows bat或cmd命令的小黑框
1. 新建文件保存为run.vbs,example.cmd 是要执行的脚本

```
Set ws = CreateObject("Wscript.Shell")    
ws.run "cmd /c C:\IIMS\serverscmd\example.cmd",0
```

参考： weixin_30723433老哥

2. 不适用vb，调用脚本，在脚本中直接指定

```
//autoStart.bat
@echo off 
if "%1" == "h" goto begin 
mshta vbscript:createobject("wscript.shell").run("""%~nx0"" h",0)(window.close)&&exit 
:begin 
//上面是添加隐藏黑框框的
//下面放自己脚本需要执行的命令
```

ps:这种在脚本中直接关闭黑框的方法，如果启动多个程序的时候，就会提示脚本错误。（注意点自己的需求就好）

参考： weixin_30648587老哥
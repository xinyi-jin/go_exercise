---
title: Git命令的简单使用
date: 2019-03-06 16:45:00
tags:
- Git
categories:
- 学习【资料整理记录】
---

- 下载Git
  windows版本下载地址：https://gitforwindows.org/
  其他版本可以去 https://git-scm.com/downloads
  下载GUI Client 图形化界面工具（比如：SourceTree）
  下载好后双击运行，我才用的傻瓜式安装，直接NEXT就好。其中有一个默认的文本编辑方式，vim（使用方法和linux一样）。

安装好后，就可以右键桌面—>Git Bash Heare 打开命令窗口了
打开之后的效果：

- 创建GitHub远程仓库
  注册GitHub账号，地址：https://github.com/
  创建Git仓库
  然后保存这里的连接用于连接远程仓库地址
  三、使用Git连接远程仓库
  打开Git命令行窗口，克隆远程仓库

```git clone https://github.com/xinyi-jin/test.git```

后边的链接指的是上一步复制的远程仓库地址
经过上边的操作，就可以把远程仓库的东西复制到本地，因为我是刚刚创建的远程仓库，还没来得及放东西，所以就提示我克隆的是一个空目录

进入克隆好的本地仓库目录，添加需要上传的文件到暂存区

```git add .```

add 后边的“.”代表本目录下所有的文件 

提交暂存区中修改的信息到本地版本库中

```git commit -m "test commit"```

-m 指的是提交信息的描述

4.最后提交本地修改到远程仓库

```git push -u origin master```

ps:这一步操作需要输入GitHub的账号密码。origin 和master 为默认信息
使用git status可以查看当前本地仓库文件的状态

四、查看GitHub远程仓库


GitHub源码：https://github.com/xinyi-jin/

ps:部分浏览器不支持，会显示日期错误
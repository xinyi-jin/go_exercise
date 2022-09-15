---
title: MySQL57的安装与使用
date: 2019-12-01 16:45:00
tags:
- MySQL
categories:
- 学习【资料整理记录】
---

- 下载安装MySQL57
1. 去官网下载MySQL57的安装包，https://dev.mysql.com/downloads/installer/
不选择最新版本，点击选择其他版本

点击下载

选择只开始我的下载

下载完成后，双击可运行程序
ps：安装过程中，能next的地方就next，这种界面就暂不提供截图分享了

设置数据库使用的端口号

设置数据库的密码

设置数据库的DOS命令启动名称

等待安装完成后Execute，然后finish就安装好了

输入用户名密码，连接到数据库，这里默认的是root最高权限账户

- DOS简单使用MySQL57
使用DOS命令操作数据库之前先要配置好MySQL的环境变量，同jdk的配置方式一样

ps：经过install安装工具连接成功之后就不用再使用下边的DOS命令，应为服务已经启动了

MySQL服务安装

```mysqld -install```

启动MySQL数据库

```net start mysql```

关闭MySQL数据库

```net stop mysql```

连接数据库（-h 数据库的IP地址，-u用户名 -p密码）

```mysql -h 127.0.0.1 -u root  -p```

ps:进入数据库之后就可以使用标准SQL语句去执行，dos下密码的输入是隐式的，就是说你输密码的时候自己看不到，只要输入后回车就可以了
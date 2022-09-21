---
title: JDK的安装与环境配置
date: 2019-02-27 16:45:00
tags:
- Java
- JDK
categories:
- 学习【资料整理记录】
---

- 下载JDK安装包
首先到Oracle官网下载，链接: https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

需要勾选同意，然后选择自己电脑系统的版本下载。
ps:右键我的电脑，属性，查看系统类型。x64是64位，x86是32位。

下载完成后，双击下载的安装包，采用傻瓜式安装，一路NEXT即可。
ps：注意安装路径可以设置在除C盘外的其他盘，例如：D:\Java\jdk-8u181-windows-x64
另外，安装路径不要使用中文命名，否则会发生未知错误（第一次装的时候就被坑死了）

- 配置JDK的环境变量
右键，我的电脑，属性，高级系统设置，高级，环境变量

找到系统变量中的path，点击编辑，在最前边添加自己的JDK路径信息，然后点击确定。
ps：路径信息到bin文件夹下就可以，因为bin文件夹存放的是java的.exe命令，路径末尾不要忘记添加“;”

然后就可以windows键+R，运行cmd，进入DOS命令行。
执行java -version命令，如下图，就表示java环境变量配置成功了。

ps:新配置好的环境变量，必须重新打开DOS窗口才能正常运行，java -version中，java后边需要留有一个空格

jdk重装时候遇到的坑0.0，没卸载 干净，这位大哥的博客有解决方案
https://blog.csdn.net/ldld1717/article/details/52144760
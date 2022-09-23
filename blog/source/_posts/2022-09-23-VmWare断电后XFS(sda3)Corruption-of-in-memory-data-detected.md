---
title: VmWare断电后XFS(sda3):Corruption of in-memory data detected
date: 2022-09-23 10:25:00
tags:
- VmWare
- 虚拟机
- 系统报错
categories:
- 学习【资料整理记录】
---

问题：VmWare断电后XFS(sda3):Corruption of in-memory data detected

原因：sda3内存数据损坏，需要进行磁盘修复操作

解决办法：

```
umount /dev/sda3
xfs_repair -v -L /dev/sda3
reboot
```

以上内容参考：[虚拟机断电后centos7无法正常启动 XFS(sda3):Corruption of in-memory data detected](https://blog.csdn.net/s19980228/article/details/107160008)

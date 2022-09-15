---
title: git ssh连接
date: 2022-09-13 12:00:00
tags:
- Git
categories:
- 学习【资料整理记录】
---

github ssh

1. // 配置git用户

```
git config --global user.email "iezhuhe@163.com"
git config --global user.name "xinyi-jin"
```

2. // genssh

```
ssh-keygen -t rsa -C "iezhuhe@163.com"
```

3. // verssh

```
ssh -T git@github.com
```


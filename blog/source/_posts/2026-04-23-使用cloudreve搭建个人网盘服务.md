---
title: 使用cloudreve搭建个人网盘服务
date: 2026-04-23 15:00:00
tags:
- Docker
- cloudreve
categories:
- 学习【资料整理记录】

---

基于有一台云服务的情况下，想玩玩云服务可以玩的项目，故有此文操作记录。

官方网址：https://cloudreve.org/

docker命令：
``
docker run -d --name cloudreve-offical \
    -p 5212:5212 \
    -p 6888:6888 \
    -p 6888:6888/udp \
    -v ~/cloudreve/data:/cloudreve/data \
    cloudreve/cloudreve:latest
``

搭建好的项目地址： http://154.58.233.77:5212/session
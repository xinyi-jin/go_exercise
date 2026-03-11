---
title: docker-hexo
date: 2026-03-10 09:00:00
tags:
- Hexo
- Docker
categories:
- 学习【资料整理记录】

---

使用docker 运行hexo 镜像
====

# 1. 删除已退出的容器
docker rm hexo-blog

# 2. 使用交互式模式运行，并保持容器运行
docker run -it -d \
  --name hexo-blog \
  -p 4000:4000 \
  -v $(pwd)/go_exercise/blog:/app \
  spurin/hexo:latest \
  sh -c "tail -f /dev/null"

# 3. 进入容器进行操作
docker exec -it hexo-blog sh

# 4. 在容器内启动 Hexo 服务
cd /app
npm install -g hexo-cli
hexo init blog
cd blog
npm install
hexo server -p 4000 -i 0.0.0.0


配置秘钥拉取博客源码
``
git config --global user.name "maomao"
git config --global user.email "ijmaomao@163.com"

ssh-keygen -t rsa -C "ijmaomao@163.com"
``

---
title: docker-etcd
date: 2026-03-10 09:00:00
tags:
- Etcd
- Docker
categories:
- 学习【资料整理记录】

---

2025.12.08
==========
docker 使用etcd容器
1. 使用etcd镜像版本 quay.io/coreos/etcd:v3.5.12 docker pull 拉取镜像
2. mkdir -p /data/containers/etcd/{data,config} 创建数据和配置目录
3. /data/containers/etcd/config/etcd.conf.yml 创建配置文件
```
name: etcd-s1
data-dir: /var/etcd
listen-client-urls: http://0.0.0.0:2379
advertise-client-urls: http://0.0.0.0:2379
listen-peer-urls: http://0.0.0.0:2380
initial-advertise-peer-urls: http://0.0.0.0:2380
initial-cluster: etcd-s1=http://0.0.0.0:2380
initial-cluster-token: etcd-cluster
initial-cluster-state: new
logger: zap
log-level: info
#log-outputs: stderr

特殊参数说明：

name： etcd member 名称，可根据实际情况修改
data-dir： etcd 数据目录，可根据实际情况修改
listen-client-urls:  client 流量监听地址，没特殊需求按文档填写即可
advertise-client-urls: 该 member 向外部通告的客户端 url 列表，单节点部署时不需要修改，集群部署模式需修改为容器所在节点对外提供服务的 IP
listen-peer-urls:  peer 流量监听地址，没特殊需求按文档填写即可
initial-advertise-peer-urls:  该 member 向同一集群内其他 member 通告的 peer url 列表，单节点部署时不需要修改，集群部署模式需修改为容器所在节点对外提供服务的 IP
initial-cluster:  初始化集群节点信息，单节点部署时不需要修改，集群部署模式需要填写集群中所有 member 的信息
initial-cluster-token:  初始化集群时使用的 token，随便写
initial-cluster-state:  初始化集群状态，可选的值为 new  或者  existing，通常采用 new
```
4. 创建 docker-compose.yml 文件
```
version: '3'

services:
  etcd:
    container_name: etcd-s1
    image: quay.io/coreos/etcd:v3.5.12
    command: /usr/local/bin/etcd --config-file=/var/lib/etcd/conf/etcd.conf.yml
    volumes:
      - ${DOCKER_VOLUME_DIRECTORY:-.}/data:/var/etcd
      - ${DOCKER_VOLUME_DIRECTORY:-.}/config/etcd.conf.yml:/var/lib/etcd/conf/etcd.conf.yml
      - "/etc/localtime:/etc/localtime:ro"
    ports:
      - 2379:2379
      - 2380:2380
    restart: always

<!-- networks:
  default:
    name: etcd-tier
    driver: bridge -->
```
5. 使用docker-compse 读取配置文件运行初始化容器
cd /data/containers/etcd
docker-compose up -d

<!-- 行不通 -->
使用etcdkeeper 操作etcd
1. 下载etcdkeeper源码
	https://codeload.github.com/evildecay/etcdkeeper/zip/refs/heads/master
2. 编译etcdkeeper启动程序
3. 运行编译好的程序
4. 浏览器访问地址

<!-- 直接使用容器运行etcdkeeper -->
1. 配置容器镜像源 阿里镜像库
```
  vim /etc/docker/daemon.json

  <!-- 添加内容 -->
  {
  "registry-mirrors": [
  "https://docker.xuanyuan.me",
  "https://xxx.mirror.aliyuncs.com"
],
  "insecure-registries": [
  "docker.xuanyuan.me"
],
  "dns": ["119.29.29.29", "114.114.114.114"]
}

ps:
bash <(wget -qO- https://xuanyuan.cloud/docker.sh)
上述命令可一键安装轩辕代理

```
3. 拉取镜像
```
  docker pull docker.io/evildecay/etcdkeeper
```
4. 创建 docker-compose.yml 文件
```
version: '3'

services:
  etcdkeeper:
    container_name: etcdkeeper-s1
    image: docker.io/evildecay/etcdkeeper:latest
    #command: /usr/local/bin/etcd --config-file=/var/lib/etcd/conf/etcd.conf.yml
    #volumes:
     # - /data/containers/etcd/data:/var/etcd
      #- /data/containers/etcd/config/etcd.conf.yml:/var/lib/etcd/conf/etcd.conf.yml
      #- "/etc/localtime:/etc/localtime:ro"
    ports:
      #- 2379:2379
      #- 2380:2380
      - 8080:8080
    restart: "no"

networks:
  default:
    driver: bridge
```
5. docker-compose up -d 启动容器
6. 打开网址访问etcdkeeper地址
  http://127.0.0.1:8080

目前etcdkeeper 设置密码时候无法访问 不清楚具体原因

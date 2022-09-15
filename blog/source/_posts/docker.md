---
title: 自学习内容
date: 2022-08-23 12:00:00
tags:
- Docker
categories:
- 学习【资料整理记录】
---

## docker

docker容器不保持任何数据
重要数据请使用外部卷存储（数据持久化）
容器可以挂载真实机目录或共享存储为卷

### 主机卷映射：

- 宿主机到docker容器映射：

```csharp
docker run -it -v /var/data:/abc myos
```

- 自定义容器内部网络

```
docker network create --subnet=172.18.0.0/16 mynet
```

- 使用自定义网络运行容器,再次重启时就不会按照启动顺序分配172.17.0.1网段的ip地址了

```
docker run -itd --name vm  --net bind --ip 172.18.0.103 centos /sbin/init
```

原文地址：https://cloud.tencent.com/developer/article/1966389

## 修复: “Error: Failed to download metadata for repo appstream” – CentOS 8系统错误

**解决方案：迁移到 CentOS Stream 8 或替代发行版**

```
dnf --disablerepo '*' --enablerepo=extras swap centos-linux-repos centos-stream-repos 
dnf distro-sync
```

docker 桥接方式映射宿主机IP端口，并挂载宿主机文件卷，初始化容器环境，赋予最高操作权限

```
docker run -itd --name goservice -p 41002:31002 -p 10989:9898 --privileged -v /root/hn_match_debug/:/root/hn_match_debug/ centos /sbin/init
```



注册mongod服务

mogod.conf

```
# ip

bind_ip=0.0.0.0

# mongodb port number

port=27017

# mongodb data file path

dbpath=/var/lib/mongo

# mongodb log file path

logpath=/var/log/mongodb/mongod.log

# mongodb automatically append log files

logappend=true
```

mongod

进入 /etc/init.d/ 目录
`# cd /etc/init.d/`
创建并编辑mongodb文件
`# vim mongodb`

```
#!/bin/bash
#chkconfig:345 61 61
#description:mongod

# mongoDB shell version v5.0.9

# connecting to: mongodb://0.0.0.0:27017

# MongoDB server version: 5.0.9

MONGO_PATH=/usr/local/mongodb5/bin/mongod
MONGO_PID=`ps -ef|grep 'mongod' | grep -v grep|awk '{print $2}'`
test -x $MONGO_PATH || exit 0

case "$1" in
  start)
     ulimit -n 2000
     echo "starting mongod server"
     $MONGO_PATH --fork --quiet -f /usr/local/mongodb5/mongod.conf
     echo "started mongod server"
     ;;
  stop)
     echo "stopping mongod server"
     if [ ! -z "$MONGO_PID" ]; then
        kill -15 $MONGO_PID
     fi
        echo "stopped mongod server"
     ;;
  status)
     ;;
  *)
     echo "usage: mongod {start|stop|status}"
     exit 1
esac

exit 0
```

上述文件配置完成后，执行一下几条命令

```
# chmod +x /etc/init.d/mongod
# chkconfig --add mongod
# chkconfig mongod on
```

启动时去除ip绑定

```
mongod --bind_ip 0.0.0.0
```

## 一：通过修改配置文件修改docker容器端口映射

1.使用**docker ps -a**命令找到要修改容器的**CONTAINER ID**

2.运行以下命令，进入该容器目录

```text
docker inspect【CONTAINER ID】| grep Id
cd /var/lib/docker/containers
```

2.停止容器

```text
docker stop [容器id]
```

3.停止主机docker服务

```text
systemctl stop docker
```

4.进入2得到的文件夹内，修改hostconfig.json 和 config.v2.json

```text
vi hostconfig.json

比如新增一个 80 端口，在PortBindings下边添加以下内容，端口配置之间用英文字符逗号隔开

"80/tcp": [ 
{
 "HostIp": "0.0.0.0",
 "HostPort": "80"
 }
]
接着修改vi config.v2.json, 找到ExposedPorts和Ports  仿照之前内容添加端口映射
"ExposedPorts":  {
    "2000/tcp":   {}
},

"Ports":{
   "2000/tcp":[
        {
         "HostIp": "0.0.0.0",
         "HostPort":  "2000"
         }
]
},
```

5.保存之后重启docker服务和容器

```text
systemctl start docker
docker start [docker id]
```

## 二：把运行中的容器生成新的镜像，再新建容器

1.提交一个运行中的容器为镜像

```text
docker commit [containerid] [new_imagename]
```

2.运行新建的镜像并添加端口映射

```text
docker run -d -p 8000:80  [imagename] /bin/sh
```

## 三：修改主机iptables端口映射

> docker的端口映射并不是在docker技术中实现的，而是通过宿主机的iptables来实现。通过控制网桥来做端口映射，类似路由器中设置路由端口映射。

如果我们有一个容器的8000端口映射到主机的9000端口，先查看iptabes设置了什么规则：

```text
sudo iptables -t nat -vnL
```

结果中有一条：

```text
Chain DOCKER (2 references)
pkts bytes target prot opt in     out     source        destination         
 98  5872 RETURN  all  --  docker0 *     0.0.0.0/0     0.0.0.0/0           
237 14316 DNAT    tcp  --  !docker0 *    0.0.0.0/0    0.0.0.0/0    tcp dpt:9000 to:172.17.0.3:8000
```

我们可以看到docker创建了一个名为DOKCER的自定义的链条Chain。而我开放8000端口的容器的ip是172.17.0.3。

也可以通过inspect命令查看容器ip

```text
docker inspect [containerId] |grep IPAddress
```

我们想再增加一个端口映射，比如`8081->81`，就在这个链条是再加一条规则：

```text
sudo iptables -t nat -A  DOCKER -p tcp --dport 8081 -j DNAT --to-destination 172.17.0.3:81
```

加错了或者想修改：先显示行号查看

```text
sudo iptables -t nat -vnL DOCKER --line-number
```

删除规则3

```text
sudo iptables -t nat -D DOCKER 3
```


---
title: linuxCmd
date: 2020-11-25 12:00:00
tags:
- 服务部署
categories:
- 学习【资料整理记录】
---

```
nohup ./game_private_server -Config=./config/game_config_mjjxhzzm_61.cfg >> log/game_private_server_mjjxhzzm_61.log 2>&1 &

nohup ./game_private_server -Config=./config/game_config_mjhnwxd_61.cfg >> log/game_private_server_mjhnwxdd_61.log 2>&1 &

nohup ./game_private_server -Config=./config/game_config_mjhnayhx_88.cfg mjhnayhx_88 >> log/game_private_server_mjhnayhx_88.log 2>&1 &
```

redis 添加测试

```
auth qwe123!@#
select 9
hset haunters id 1，
HEXISTS haunters
```

Redis 设置密码

```
  config set requirepass password
  auth password
  config get requirepass
```




```
cmd:
  config set requirepass qwe123!@#
  auth qwe123!@#
  config get requirepass
```

optional string PassWord = 3;		// 密码

  ./protoc.exe --go_out=. ./*.proto

  "notice#body"

  624691


-- mysql开启事件命令

```
show variables like '%sche%';

set global event_scheduler =1
```

天天 userid 406570

// ssh连接服务器

```
sudo ssh root@39.100.68.248
```

// 上传文件到服务器

```
sudo scp /Users/xinyi-jin/Desktop/test.txt root@39.100.68.248:/root/server/
```

// 上传文件到服务器，不会影响当前正在运行的程序

```
sudo rsync -av /Users/xinyi-jin/go/src/jxserver/bin/xiangque/* root@39.100.68.248:/root/server
```

```
cygwin错误 
原因：ssh.exe可能会与git下边的版本冲突
解决办法：可使用-e参数指定ssh.exe

rsync -av -e D:\soft\cwrsync_6.2.4_x64_free\bin\ssh.exe robot/robot root@47.107.37.120:/root/%1/robot/
```

// mac 交叉编译linux

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build gripper.go
```


Gitshell中输入如下命令解决：

```
git config --global core.autocrlf false
```
Go 自动生成测试文件
```
gotests -only IntToRoman -w intToRoman.go
```

WinSCP命令

```
./WinSCP.com   进入winscp命令行
open sftp://root@47.107.37.120 -privatekey=C:\Users\Administrator\.ssh\id_rsa.ppk   连接远端
```


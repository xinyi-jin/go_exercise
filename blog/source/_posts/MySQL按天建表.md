---
title: MySQL按天建表
date: 2020-01-02 16:45:00
tags:
- MySQL
categories:
- 学习【资料整理记录】
---

mysql设置定时任务：

需求：mysql数据库里面做个定时任务自动每天创建一张表。

1. 查看是否开启event与开启event

```show variables like '%sche%';  ```

若未开启

```set global event_scheduler =1;```

2. 创建存储过程

每天创建一个名为 test_年月日的表


```
delimiter //

drop procedure if exists create_table01//

CREATE PROCEDURE create_table01()

BEGIN

declare str_date varchar(50);

SET str_date = date_format(now(),'%Y%m%d');  

 

SET @sqlcmd1 = CONCAT('CREATE TABLE test_',str_date,'(

id int(11) not null auto_increment primary key,

time datetime not null

) ;');

PREPARE p1 FROM @sqlcmd1;

EXECUTE p1;

DEALLOCATE PREPARE p1;

END//

delimiter ;
```

3. 创建事件，调用存储过程

```
drop event if exists test_event01;

create event test_event01

on schedule every 1 DAY STARTS NOW()

do call create_table01();
```

开启事件

```alter event test_event on completion preserve enable;```

```show tables;```

查询是否建表成功

若是想关闭事件：

```alter event test_event on completion preserve disable;```

4. 事件执行时间的设置

```
CREATE EVENT test_event01 ON SCHEDULE EVERY 1 DAY STARTS '2012-09-24 00:00:00'

ON COMPLETION PRESERVE ENABLE DO CALL  create_table01();

EVERY #后面的是时间间隔，可以选 1 second，3 minute，5 hour，9 day，1 month，1 quarter（季度），1 year

#从2013年1月13号0点开始，每天运行一次

ON SCHEDULE EVERY 1 DAY   STARTS '2013-01-13 00:00:00'

#从现在开始每隔九天定时执行

ON SCHEDULE EVERY 9 DAY STARTS NOW() ；

#每个月的一号凌晨1 点执行

on schedule every 1 month starts date_add(date_add(date_sub(curdate(),interval day(curdate())-1 day),interval 1 month),interval 1 hour);

#每个季度一号的凌晨1点执行

on schedule every 1 quarter starts date_add(date_add(date(concat(year(curdate()),'-',elt(quarter(curdate()),1,4,7,10),'-',1)),interval 1 quarter),interval 1 hour);

#每年1月1号凌晨1点执行

on schedule every 1 quarter starts date_add(date_add(date(concat(year(curdate()),'-',elt(quarter(curdate()),1,4,7,10),'-',1)),interval 1 quarter),interval 1 hour);
```

5. 几个相关概念

存储程序：( 存储函数(stored function)+存储过程(stored procedure)+触发器(trigger)+事件(event) )

存储函数：返回一个计算结果，该结果可以用在表达式里

存储过程：不直接返回一个结果，但可以用来完成一般的运算或是生成一个结果集并传递会客户

触发器：与数据表关联，当那个数据表被 insert、delete、update语句修改时，触发器将自动执行

事件：根据时间表在预订时刻自动执行

原文地址：https://blog.csdn.net/vinking9393/article/details/84805050
---
title: MySQL中WHERE子句不等于的使用
date: 2020-03-17 16:45:00
tags:
- MySQL
categories:
- 学习【资料整理记录】
---

MySQL中WHERE子句条件判断<>，!= 二者的区别
ANSI标准中是用<>(所以建议用<>)，但为了跟大部分数据库保持一致，数据库中一般都提供了 !=(高级语言一般用来表示不等于) 与 <> 来表示不等于：

- MySQL 5.1: 支持 != 和 <>
- PostgreSQL 8.3: 支持 != 和 <>
- SQLite: 支持 != 和 <>
- Oracle 10g: 支持 != 和 <>
- Microsoft SQL Server 2000/2005/2008: 支持 != 和 <>
- IBM Informix Dynamic Server 10: 支持 != 和 <>
- InterBase/Firebird: 支持 != 和 <>

最后两个只支持ANSI标准的数据库：
- IBM DB2 UDB 9.5:仅支持 <>
- Apache Derby:仅支持 <>

建议使用<>标识不等于关系，这样适配所有的数据库
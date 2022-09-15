---
title: JS两个日期之间计算时间差
date: 2019-03-06 16:45:00
tags:
- Html
- Js
categories:
- 学习【资料整理记录】
---

- 计算两个日期之间相差的毫秒数（也就是两个日期的时间戳差）
这里有几种常见的JS方式，来获得日期的毫秒数。

1. 日期对象的 getTime();方法
2. 
```
<script>
    var date = new Date(); //新建一个日期对象，默认现在的时间
	var timestamp = date.getTime();	//调用getTime()方法获取毫秒数
	alert("timestamp: "+timestamp);
</script>
```

2. 日期对象的valueOf();方法
3. 
```
<script>
    var date = new Date(); //新建一个日期对象，默认现在的时间
	var timestamp = date.valueOf(date); //调用valueOf()方法获取毫秒数
	alert("timestamp: "+timestamp);
</script>
```
3. Date对象的parse();方法

```
<script>
    var date = new Date(); //新建一个日期对象，默认现在的时间
	var timestamp = Date.parse(date);//使用Date对象的parse()方法，获取毫秒数
	alert("timestamp: "+timestamp);
</script>
```

- 用现在时间的毫秒数减去过去时间的毫秒数，得到的就是两个日期相差的总毫秒数。
ps：可以使用上述的三种方法，另外还有一种最为简单的方法。如果两个日期对象都是Date类型的话，可以直接相减，得到的就是毫秒数差。

```
<script>
	var new_date = new Date(); //新建一个日期对象，默认现在的时间
	var old_date = new Date("2018-12-12 00:00:00"); //设置过去的一个时间点，"yyyy-MM-dd HH:mm:ss"格式化日期

	var difftime = new_date - old_date; //计算时间差
	
	alert("difftime: "+difftime);
</script>
```

- 得到了两个日期的毫秒差之后，就可以转化成具体的日期格式（xxxx年xx月xx天，xx小时xx分钟xx秒）
因为我们获取的是毫秒，所有要先转换成秒。1秒=1000毫秒

```
<script>
	var new_date = new Date(); //新建一个日期对象，默认现在的时间
	var old_date = new Date("2018-12-12 00:00:00"); //设置过去的一个时间点，"yyyy-MM-dd HH:mm:ss"格式化日期

	var difftime = (new_date - old_date)/1000; //计算时间差,并把毫秒转换成秒
	
	var days = parseInt(difftime/86400); // 天  24*60*60*1000 
	var hours = parseInt(difftime/3600)-24*days;    // 小时 60*60 总小时数-过去的小时数=现在的小时数 
	var minutes = parseInt(difftime%3600/60); // 分钟 -(day*24) 以60秒为一整份 取余 剩下秒数 秒数/60 就是分钟数
	var seconds = parseInt(difftime%60);  // 以60秒为一整份 取余 剩下秒数
	
	alert("时间差是: "+days+"天, "+hours+"小时, "+minutes+"分钟, "+seconds+"秒");		
</script>
```

另外本人踩过坑的一个项目，还是从网上获取的别人的代码，不曾想还是个有bug的项目。
1
当时这种方式设置日期，然后再获取毫秒数，就会出现30天的时间差。后来干脆直接格式化一下日期放进去，没想到就对了。我也很懵比，有知道的大牛，可以告知一下这样做为什么会好，万分感谢！！！

GitHub源码：https://github.com/xinyi-jin/Romantic-confession

ps:部分浏览器不支持，会显示日期错误
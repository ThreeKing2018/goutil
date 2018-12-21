# golang 时间函数,快捷获取

> 做项目离不开跟时间打交道,如果有快捷的小函数使用,何乐而不为呢

# 安装
```
go get -u github.com/yezihack/gotime
```


# 使用方法
```golang
t := gotime.New()
//获取当前时间
tm := t.Now()

# 时间区
* 默认中国东八区

# 使用方法
```golang
t := gotime.New()
//获取当前时间
tm := t.Now()
fmt.Println(tm)
//print 2018-12-12 15:16:19

//获取当前零晨时间
tm1 := gotime.New().NowStart()
fmt.Println(tm1)
//print 2018-12-12 00:00:00

//获取一个最后一刻
tm2 := gotime.New().NowEnd()
fmt.Println(tm2)
//print 2018-12-12 23:59:59

//获取当前时间戳
tm3 := gotime.New().NowUnix()
fmt.Println(tm3)
//1544599059


```

# 万能的format转换格式函数
```
//自定义格式的万能转换时间函数
tm := gotime.New().Format(gotime.New().NowTime(), "YYYY-MM-DD HH:mm:ss")
fmt.Println(tm)

//时间
tm1 := gotime.New().Format(gotime.New().NowTime(), "YYYY-MM-DD")
fmt.Println(tm1)

tm2 := gotime.New().Format(gotime.New().NowTime(), "HH:mm:ss")
fmt.Println(tm2)

tm3 := gotime.New().Format(gotime.New().NowTime(), "YYYY/MM/DD HH:mm:ss")
fmt.Println(tm3)
```

# 查询列表

| 函数名称 | 说明 | 实例 |
| ---: | ---:| ---:|
|Now| 获取当前时间 | 2018-12-12 15:22:22|
|NowUnix| 获取当前时间戳 | 2018-12-12 15:22:22|
|NowTime| 获取当前时间 time格式 | 2018-12-12 15:21:53.956340732 +0800 Asia/Shanghai |
|GetYmd| 获取当前年月日 | 2018-12-12 |
|GetHms| 获取当前分时秒 | 15:22:22|
|NowStart| 获取今天零晨时间 | 2018-12-12 00:00:00|
|NowEnd| 获取今天最后一秒 | 2018-12-12 23:59:59|
|Before| 时间相减,返回字符串 | 2018-12-12 15:22:22|
|Next| 时间相加,返回字符串 | 2018-12-12 15:22:22|
|RfcToUnix| RFC格式转时间戳 | 1544599696|
|GetRFC3339| 获取RFC格式 | 2018-12-12T15:32:25+08:00|
|ToRFC3339| 转换成RFC格式 | 2018-12-12T15:32:25+08:00|
|Format| 转换成任意格式 | |
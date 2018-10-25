# goutil
> 是go 的工具集,也叫小函数 <br/>打造一站式工具库<br/>会持续更新

## 目录说明
- 一个类别一个文件夹,文件夹里必含一个测试文件
- 可以使用go test测试
- 如array

```
 |- array
        |- array.go      //核心文件
        |- array_test.go //测试文件
```

## 目录管理
- 采用govendor包管理

## 目录分类列表

| 功能 | 包名 |  备注 |
| :--- | :--- | :--- |
| 数值转换 | [convertor](convertor) | 操作数字等 |
| 时间操作 | [datetime](datetime) |  获取自定义时间格式等|
| 数组操作 | [array](array) |  数组转换等|
| 文件操作 | [filetool](filetool) |  获取文件目录,读取,写等|
| 格式化操作 | [formatter](formatter) |  如存储大小转换成可读的单位等|
| 日志操作 | [logtool](logtool) |  打印不同等级日志等|
| 分页操作 | [paginator](paginator) |  用于数据分页操作等|
| rpc操作 | [rpctool](rpctool) |  rpc等|
| 切片操作 | [slicetool](slicetool) |  切片操作等|
| 字符串操作 | [strtool](strtool) |  随机数,md5等|
| 命令操作 | [slicetool](slicetool) |  linux相关等|


## 贡献来源
> 一般来源于github和我们自己写的

### 部分来源名单

- https://github.com/UlricQin/goutils
- https://github.com/henrylee2cn/goutil/
- https://github.com/wudaoluo/goutil

# 目录
- [前言](preface.md)

- gotime
    对系统时间进行的操作
- golog

    参考 https://github.com/xiaomeng79/go-log.git
    模块化的日志

- goaddr
    获取内网地址和外网地址

- config

    参考：
        https://github.com/spf13/viper.git
        https://github.com/kelseyhightower/confd.git
    配置文件-支持本地+远程配置文件读取,动态加载

    TODO: 暂时支持etcdv2 ,etcdv3在近期添加和yaml支持


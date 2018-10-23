# goutil
是go 的工具集,也叫小函数

## 目录要求
- 一个类别一个文件夹,文件夹里必含一个测试文件
- 如array

```
 |- array
        |- array.go      //核心文件
        |- array_test.go //测试文件
```

## 来源
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


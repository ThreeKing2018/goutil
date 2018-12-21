# goutil
> 是go 的工具集,也叫小函数 <br/>打造一站式工具库<br/>会持续更新

# 安装
```
go get -u github.com/ThreeKing2018/goutil
```


## 目录说明
- 一个类别一个文件夹,文件夹里必含一个测试文件
- 可以使用go test测试
- 如array

```
 |- array
        |- array.go      //核心文件
        |- array_test.go //测试文件
        |- readme.md     //说明文档
```

## 目录分类列表

| 功能 | 包名 |  备注 |
| :--- | :--- | :--- |
| 目录操作 | [pwdtools](pwdtools/pwdtools.go) | 获取目录 |
| 数值转换 | [convertor](convertor/readme.md) | 操作数字等 |
| 时间操作 | [time](time/readme.md) |  获取自定义时间格式等|
| 数组操作 | [array](array) |  数组转换等|
| 文件操作 | [filetool](filetool) |  获取文件目录,读取,写等|
| 格式化操作 | [formatter](formatter) |  如存储大小转换成可读的单位等|
| 日志操作 | [logtool](logtool) | 简单好用, 可以打印不同等级日志等|
| golog | [golog](golog) |  操作日志等|
| 分页操作 | [paginator](paginator) |  用于数据分页操作等|
| rpc操作 | [rpctool](rpctool) |  rpc等|
| 切片操作 | [slicetool](slicetool) |  切片操作等|
| 字符串操作 | [strtool](strtool) |  随机数,md5等|
| 命令操作 | [slicetool](slicetool) |  linux相关等|
| goaddr | [goaddr](goaddr/readme.md) |  获取内网地址和外网地址|
| config | [config](config/readme.md) |  配置文件-支持本地+远程配置文件读取,动态加载|
| grpc+etcd服务注册 | [register](register/README.md) |  grpc的resolver方式服务注册|
| 常用hash函数 | [hash](hash/README.md) |  string、byte、file 的hash值 包括md5 sha1 sha256 sha512 |
| curl | [curl](curl/curl.go) |  curl get ,post 请求 |
| 第三方免费服务 | [Three_service](three_service/bank.go) |  如: 在线验证银行卡|
| 定时器 | [Timer](timer/timer.go) |  原生,实现一个简单的定时器 |
| 实现阻塞 | [choke](choke/README.md) | 一般用于程序阻塞 , 简单好用 |


## 贡献来源
> 一般来源于github和我们自己写的

### 部分来源名单

- https://github.com/UlricQin/goutils
- https://github.com/henrylee2cn/goutil/
- https://github.com/wudaoluo/goutil


## 欢迎加入本团队
> 我们正在打造一个一站式工具库<br/>实现快速开发,做到开箱即用<br/>方便大家,请多多支持,加个星吧



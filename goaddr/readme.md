# 使用方式

### 获取内网地址
    NewAddr(4440).IntranetAddr().GetIPstr()

### 获取外网地址
     NewAddr(4441).ExternalAddr().GetIPstr()

### 获取内网地址 排除ip 172.16.21.77
    NewAddr(4442).Exclude([]string{"172.16.21.77"}).GetIPstr()
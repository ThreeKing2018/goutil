## register grpc的resolver方式服务注册

- 由于grpc官方不在建议使用 `google.golang.org/grpc/naming` 进行服务注册，
- 建议使用 `google.golang.org/grpc/resolver` 包进行服务注册。
- 此处是服务注册的一个实现代码

## 备注

来源：[http://morecrazy.github.io/2018/08/14/grpc-go%E5%9F%BA%E4%BA%8Eetcd%E5%AE%9E%E7%8E%B0%E6%9C%8D%E5%8A%A1%E5%8F%91%E7%8E%B0%E6%9C%BA%E5%88%B6/](http://morecrazy.github.io/2018/08/14/grpc-go%E5%9F%BA%E4%BA%8Eetcd%E5%AE%9E%E7%8E%B0%E6%9C%8D%E5%8A%A1%E5%8F%91%E7%8E%B0%E6%9C%BA%E5%88%B6/)
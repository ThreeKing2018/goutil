package register

import (
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	etcdAddrs  = "127.0.0.1:2379"  // 多个使用';'分开
	svcName    = "svc"             // 微服务名称
	listenAddr = "127.0.0.1:28080" // grpc服务监听地址 - 本test未实现grpc监听，测试时请填写正确的grpc监听地址
)

// Test_RegisterServer 测试服务注册
func Test_RegisterService(t *testing.T) {
	// 注册服务
	err := Register(etcdAddrs, svcName, listenAddr, 5)
	if err != nil {
		t.Errorf("服务注册失败:%v", err)
	}
}

// 测试客户端发现
func Test_DiscoveryService(t *testing.T) {
	r := NewResolver(etcdAddrs)
	resolver.Register(r)
	t.Logf(r.Scheme() + "://" + svcName)
	_, err := grpc.Dial(r.Scheme()+"://author/"+svcName, grpc.WithBalancerName("round_robin"), grpc.WithInsecure(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		t.Errorf("连接grpc服务失败:%v", err)
	}
}

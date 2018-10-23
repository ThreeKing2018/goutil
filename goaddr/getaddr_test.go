package getaddr

import (
	"fmt"
	"testing"
)

//str
func Test_GetneiIp_str(t *testing.T) {
	ipaddr := NewAddr(4440).IntranetAddr().GetIPstr()
	fmt.Println("内网地址", ipaddr)
}

func Test_GetwaiIp_str(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Test_GetwaiIp_str没有找到外网地址")
		}
	}()

	ipaddr := NewAddr(4441).ExternalAddr().GetIPstr()
	fmt.Println("外网地址", ipaddr)
}

func Test_ExcludeIp_str(t *testing.T) {
	ipaddr := NewAddr(4442).Exclude([]string{"172.16.21.77"}).GetIPstr()
	fmt.Println("ExcludeIp", ipaddr)

}

func Test_Getlocal_str(t *testing.T) {
	a := NewAddr(4002).LocalAddr().GetIPstr()

	fmt.Println("本地地址", a)
}

//tcp
func Test_GetneiIp_tcp(t *testing.T) {
	a := NewAddr(4440)
	a.IntranetAddr()
	fmt.Println(a.GetTCPAddr())
}

func Test_GetwaiIp_tcp(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Test_GetwaiIp_tcp没有找到外网地址")
		}
	}()
	a := NewAddr(4441)
	a.ExternalAddr()
	fmt.Println(a.GetTCPAddr())
}

//udp
func Test_GetneiIp_udp(t *testing.T) {
	a := NewAddr(4440)
	a.IntranetAddr()
	fmt.Println(a.GetUDPAddr())
}

func Test_GetwaiIp_udp(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Test_GetwaiIp_udp没有找到外网地址")
		}
	}()
	a := NewAddr(4441)
	a.ExternalAddr()
	fmt.Println(a.GetUDPAddr())
}

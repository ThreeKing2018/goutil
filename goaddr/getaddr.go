package getaddr

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//在mac和linux测试可用  windows没有测试过

const UDPaddr = "223.5.5.5:53" //udp没有连接的概念所以这个地址随便写一个就行了

type Getaddr interface {
}

type getaddr struct {
	list        []net.Addr
	ip          net.IP
	port        int
	excludeList []string
}

//_type {false:内网地址 ,true:外网地址}
func NewAddr(port int) *getaddr {
	return &getaddr{
		port: port,
	}
}

func (ip *getaddr) Exclude(iplist []string) *getaddr {
	ip.excludeList = iplist
	return ip
}

func (ip *getaddr) LocalAddr() *getaddr {
	ip.ip = net.ParseIP("127.0.0.1")
	return ip
}

func (ip *getaddr) IntranetAddr() *getaddr {
	_ip := ip.getip(false)
	if _ip == nil {
		panic("没有找到内网地址")
		//return fmt.Errorf("没有找到内网地址")
	}
	ip.ip = _ip
	return ip
}

func (ip *getaddr) ExternalAddr() *getaddr {
	_ip := ip.getip(true)
	if _ip == nil {
		panic("没有找到外网地址")
		//return fmt.Errorf("没有找到外网地址")
	}
	ip.ip = _ip
	return ip
}

func (addr *getaddr) GetIPstr() string {
	return fmt.Sprintf("%s:%d", addr.ip.String(), addr.port)
}

func (addr *getaddr) GetTCPAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: addr.ip, Port: addr.port}
}

func (addr *getaddr) GetUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{IP: addr.ip, Port: addr.port}
}

func isPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

func getPulicIP() net.IP {
	conn, _ := net.DialTimeout("udp", UDPaddr, 1*time.Second)

	idx := strings.Split(conn.LocalAddr().String(), ":")[0]
	//关闭连接
	conn.Close()
	_ip := net.ParseIP(idx)
	if _ip != nil {

		return _ip
	}

	//logger.Fatal("获取外网地址失败",idx)
	return nil
}

func (addr *getaddr) getip(_type bool) net.IP {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
		//logger.Fatal("[bug] 获取地址列表错误", err)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		ipnet, ok := address.(*net.IPNet) //断言
		if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			for _, addr := range addr.excludeList {
				if ipnet.String() == addr {
					continue
				}
			}
			if _type && isPublicIP(ipnet.IP) { //获取外网地址
				//logger.Info("找到外网地址",ipnet.IP.String())
				return ipnet.IP

			} else if !_type && !isPublicIP(ipnet.IP) { //获取内网地址
				//logger.Info("找到内网地址",ipnet.IP.String())
				return ipnet.IP
			}

		}

	}

	//获取通过nat转换上网的服务器地址
	//logger.Info("通过ip没有找到外网地址，开始联网分析")
	if _type {
		_ip := getPulicIP()
		if _ip != nil {
			return nil
			//logger.Info("联网找到能够上网的地址，这个地址是内网类型经过nat转换的",_ip.String())
		}
		return _ip
	}

	return nil
}

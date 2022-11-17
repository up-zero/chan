package define

import (
	"strconv"
	"strings"
)

// RoundRobin 轮询，加权轮询
func (u *Upstream) RoundRobin(proxyPass string) string {
	beginIndex := 0
	index := 0
	for i, value := range u.Values {
		endIndex := beginIndex + value.Weight
		if beginIndex <= u.Counter && u.Counter < endIndex {
			index = i
			break
		}
		beginIndex = endIndex
	}
	value := u.Values[index].Value
	u.Counter = (u.Counter + 1) % u.Weights
	return strings.ReplaceAll(proxyPass, u.Name, value)
}

// IpHash 根据客户端的IP，计算负载地址
func (u *Upstream) IpHash(proxyPass, remoteAddr string) string {
	clientIp := ClientIP(remoteAddr)
	// 本地默认返回0号负载地址
	if len(u.Values) > 0 && clientIp == "" {
		return strings.ReplaceAll(proxyPass, u.Name, u.Values[0].Value)
	}
	// 计算ip映射的权重位置
	addr := strings.Split(clientIp, ".")
	count := 0
	for i := 0; i < len(addr) && i < 3; i++ {
		addrInt, _ := strconv.Atoi(addr[i])
		count += addrInt
	}
	count %= u.Weights
	// 计算权重对应的落点位置
	beginIndex := 0
	index := 0
	for i, value := range u.Values {
		endIndex := beginIndex + value.Weight
		if beginIndex <= count && count < endIndex {
			index = i
			break
		}
		beginIndex = endIndex
	}
	return strings.ReplaceAll(proxyPass, u.Name, u.Values[index].Value)
}

package define

import "net"

func ClientIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return ""
	}
	remoteIp := net.ParseIP(ip)
	if remoteIp == nil {
		return ""
	}
	return remoteIp.String()
}

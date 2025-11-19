package util

import (
	"context"
	"encoding/binary"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func GenerateSpanID(addr string) string {
	times := uint64(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	// 尝试解析IP地址，如果失败则使用默认值
	var ipLong uint32 = 0x7f000001 // 默认127.0.0.1

	// 支持IPv4和IPv6的安全解析
	if host, _, err := net.SplitHostPort(addr); err == nil {
		if ip := net.ParseIP(host); ip != nil {
			if ipv4 := ip.To4(); ipv4 != nil {
				ipLong = binary.BigEndian.Uint32(ipv4)
			}
			// IPv6地址保持使用默认值，避免复杂的哈希计算
		}
	}
	spanId := ((times ^ uint64(ipLong)) << 32) | uint64(rand.Int31())
	return strconv.FormatUint(spanId, 16)
}

func Ip2Long(ip string) (uint32, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(ipAddr.IP.To4()), nil
}

func GetTraceInfoFromCtx(ctx context.Context) (traceId, spanId, pSpanId string) {
	if ctx.Value("traceid") != nil {
		traceId = ctx.Value("traceid").(string)
	}
	if ctx.Value("spanid") != nil {
		spanId = ctx.Value("spanid").(string)
	}
	if ctx.Value("pspanid") != nil {
		pSpanId = ctx.Value("pspanid").(string)
	}

	return
}

// GetLocalIP 获取本机的第一个非回环IP地址
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

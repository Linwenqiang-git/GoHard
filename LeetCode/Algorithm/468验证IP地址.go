package algorithm

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	ret := validIPV4Address(queryIP)
	if ret == "Neither" {
		ret = validIPV6Address(queryIP)
	}
	return ret
}
func validIPV4Address(queryIP string) string {
	ipSegment := strings.Split(queryIP, ".")
	if len(ipSegment) == 4 {
		for _, s := range ipSegment {
			if len(s) > 1 && s[0] == '0' {
				return "Neither"
			}
			intvalue, err := strconv.Atoi(s)
			if err != nil || intvalue > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	return "Neither"
}

func validIPV6Address(queryIP string) string {
	ipSegment := strings.Split(queryIP, ":")
	if len(ipSegment) == 8 {
		for _, value := range ipSegment {
			if value == "" || len(value) > 4 || len(value) < 1 {
				return "Neither"
			}
			_, err := strconv.ParseUint(value, 16, 64)
			if err != nil {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

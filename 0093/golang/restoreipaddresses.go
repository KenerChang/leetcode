package restoreipaddresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	// 1. ip address should have 4 integer
	// 2. ip address is between 0 ~ 255
	// 3. no single digit left

	cache := map[string][][]string{}
	ips := restoreIpAddressesImpl(s, 4, cache)

	results := []string{}
	for _, ip := range ips {
		results = append(results, strings.Join(ip, "."))
	}
	return results
}

func restoreIpAddressesImpl(s string, remain int, cache map[string][][]string) (result [][]string) {
	sLen := len(s)
	if sLen > 12 {
		return
	}

	if s == "" {
		return
	}

	if remain == 1 {
		if s[0] == '0' && sLen > 1 {
			return
		}

		intS, _ := strconv.Atoi(s)
		if intS <= 255 && s != "" {
			return [][]string{{s}}
		}
		return
	}

	key := s + "x" + strconv.Itoa(remain)
	if ips, found := cache[key]; found {
		return ips
	}

	i := 1
	for i < 4 {
		if sLen < i {
			break
		}

		if s[0] == '0' && i > 1 {
			break
		}

		subStr := s[0:i]
		num, _ := strconv.Atoi(subStr)
		if num > 255 {
			break
		}

		ips := restoreIpAddressesImpl(s[i:sLen], remain-1, cache)
		i++
		if len(ips) == 0 {
			continue
		}

		for _, ip := range ips {
			ipSnippet := append([]string{subStr}, ip...)
			result = append(result, ipSnippet)
		}
	}

	return
}

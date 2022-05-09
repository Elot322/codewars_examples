package kata

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	arrStr := strings.Split(ip, ".")

	if len(arrStr) == 4 {
		for _, v := range arrStr {
			val, _ := strconv.Atoi(v)
			if v == strconv.Itoa(val) && val <= 255 && val >= 0 {
				continue
			} else {
				return false
			}
		}
		return true
	}

	return false
}

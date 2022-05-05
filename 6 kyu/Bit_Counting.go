package kata

import (
	"strconv"
)

func CountBits(val uint) int {
	count := 0
	s := strconv.Itoa(int(val))
	value, _ := ConvertInt(s, 10, 2)
	for _, v := range value {
		if string(v) == "1" {
			count++
		}
	}
	return count
}

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}

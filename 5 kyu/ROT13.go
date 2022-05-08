package kata

import "strings"

func Rot13(msg string) string {
	mapped := strings.Map(converted, msg)
	return mapped
}

func converted(r rune) rune {
	if r >= 'a' && r <= 'z' {
		if r > 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		if r > 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}

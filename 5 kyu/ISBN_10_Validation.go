package kata

import "regexp"

func ValidISBN10(isbn string) bool {
	var result int
	if len(isbn) != 10 || regexp.MustCompile(`^[a-wA-W]+$`).MatchString(isbn) {
		return false
	} else {
		for i, v := range isbn {
			if v == 'X' {
				if i == len(isbn)-1 {
					result += (i + 1) * 10
				} else {
					return false
				}
			} else {
				result += (i + 1) * int(v-'0')
			}
		}
		if result%11 == 0 {
			return true
		} else {
			return false
		}
	}
}

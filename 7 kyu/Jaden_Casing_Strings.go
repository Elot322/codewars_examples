package kata

import "strings"

func ToJadenCase(str string) string {
	var resultStrings []string
	words := strings.Fields(str)
	for _, word := range words {
		var resultWord string
		chars := []rune(word)
		for i, v := range chars {
			if i == 0 {
				if v >= 97 && v <= 122 {
					v -= 32
					resultWord += string(v)
				} else {
					resultWord += string(v)
				}
			} else {
				resultWord += string(v)
			}
		}
		resultStrings = append(resultStrings, resultWord)
	}
	return strings.Join(resultStrings, " ")
}

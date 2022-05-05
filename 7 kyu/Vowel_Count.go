package kata

func GetCount(str string) (count int) {
	for _, v := range str {
		if string(v) == "a" || string(v) == "e" || string(v) == "i" || string(v) == "o" || string(v) == "u" {
			count++
		}
	}
	return count
}

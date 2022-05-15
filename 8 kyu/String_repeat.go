package kata

import "strings"

func RepeatStr(repetitions int, value string) string {
	sb := strings.Builder{}
	for i := 1; i <= repetitions; i++ {
		sb.WriteString(value)
	}
	return sb.String()
}

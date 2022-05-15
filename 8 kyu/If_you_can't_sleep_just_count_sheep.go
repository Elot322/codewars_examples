package kata

import (
	"strconv"
	"strings"
)

func countSheep(num int) string {
	sb := strings.Builder{}
	for i := 1; i <= num; i++ {
		sb.WriteString(strconv.Itoa(i) + " sheep...")
	}
	return sb.String()
}

package kata

import "strings"

func DNAtoRNA(dna string) string {
	sb := strings.Builder{}
	for _, v := range dna {
		if string(v) == "T" {
			sb.WriteString("U")

		} else {
			sb.WriteString(string(v))
		}
	}
	return sb.String()
}

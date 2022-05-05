package kata

import (
	"strings"
)

func toWeirdCase(str string) string {
	var s2 []string

	sSplit := strings.Split(str, " ")
	var sOutSplit []string
	var finalString []string

	for _, v := range sSplit {
		chars := []rune(v)
		for i, v := range chars {
			if i%2 == 0 {
				s2 = append(s2, strings.ToUpper(string(v)))
			} else {
				s2 = append(s2, strings.ToLower(string(v)))
			}
		}
		sOutSplit = append(sOutSplit, s2...)
		s3 := strings.Join(sOutSplit, "")
		sOutSplit = nil
		s2 = nil
		finalString = append(finalString, s3)
	}
	s4 := strings.Join(finalString, " ")
	return s4
}

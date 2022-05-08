package kata

import (
	"fmt"
	"sort"
	"strconv"
)

//TODO: MAKE REFACTOR! Execution Timed Out (12000 ms)
type EMPTY struct{}

var PRESENT EMPTY = EMPTY{}

func partition(val int) []int {
	var arr []int
	set := make(map[string]EMPTY)
	for i := 1; i <= val; i++ {
		for _, result := range worker([]int{i}, val-i) {
			if len(result)*len(set) < 120 {
				set[fmt.Sprintf("%v", result)] = PRESENT
			}

		}
	}
	out := make([]string, len(set))
	pos := 0
	for key := range set {
		out[pos] = key
		pos++
	}
	sort.Strings(out)
	for _, v := range out {
		value, _ := strconv.Atoi(v)
		arr = append(arr, value)
	}
	return arr
}

func worker(already []int, val int) [][]int {
	if val == 0 {
		sort.Ints(already)
		return [][]int{already}
	}
	out := [][]int{}
	for i := 1; i <= val; i++ {
		out = append(out, worker(append(already, i), val-i)...)
	}

	return out
}

func Partitions(n int) int {
	return len(partition(n))
}

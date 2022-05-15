package kata

func PositiveSum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}
	return sum
}

package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	var count int
	var sum int

	for _, v := range numbers {
		if v > 0 {
			count++
		} else {
			sum += v
		}
	}
	res = append(res, count)
	res = append(res, sum)

	return res // your code here
}

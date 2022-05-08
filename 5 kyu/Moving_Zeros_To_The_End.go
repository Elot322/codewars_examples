package kata

func MoveZeros(arr []int) []int {

	var resultArr []int
	var count int

	for _, v := range arr {
		if v == 0 {
			count++
		} else {
			resultArr = append(resultArr, v)
		}
	}

	for i := 1; i <= count; i++ {
		resultArr = append(resultArr, 0)
	}

	return resultArr
}

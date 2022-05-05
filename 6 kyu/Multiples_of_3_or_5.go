package kata

func Multiple3And5(number int) int {
	var sum int
	for i := 0; i < number; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}

	if sum < 0 {
		return 0
	} else {
		return sum
	}
}

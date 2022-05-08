package kata

func DblLinear(n int) int {
	result := []int{1}
	x, y := 0, 0
	for i := 0; i < n; i++ {
		tempX := result[x]*2 + 1
		tempY := result[y]*3 + 1
		if tempX <= tempY {
			x++
			result = append(result, tempX)
			if tempY == tempX {
				y++
			}
		} else {
			y++
			result = append(result, tempY)

		}
	}
	return result[n]
}

package normal

func Max(arr []int) int {
	max := 0
	tempmax := 0
	for _, temp := range arr {
		tempmax += temp
		if tempmax > max {
			max = tempmax
		}
		if tempmax < 0 {
			tempmax = 0
		}
	}
	return max
}

package main

import "fmt"

func main()  {
	max()
}

func max()  {
	arr := []int{-1, 2, 5, -7, -3, 0}
	max := 0
	tempmax := 0
	for _, temp := range arr {
		tempmax = tempmax + temp
		if tempmax > max {
			max = tempmax
		}
		if tempmax < 0 {
			tempmax = 0
		}
	}
	fmt.Println(max)
}
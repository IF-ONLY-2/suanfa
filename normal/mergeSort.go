package main

import "fmt"

func main() {
	arr := []int{9,8,7,6,5,4,3,2,1}

	res := mergeSort(arr)

	fmt.Println(res)
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int,0,len(left)+len(right))
	i, j := 0, 0
	for  i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
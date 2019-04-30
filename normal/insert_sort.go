package main

import "fmt"

func main() {
	arr := []int{3,1,2,4,6,7,8}

	insertSort(arr)

	fmt.Println(arr)
}

func insertSort(arr []int) {
	for i:=1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1] ; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
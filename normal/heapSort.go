package main

import "fmt"

func main() {
	arr := []int{55,2,6,7,12,34,77,89}
	//构造堆
	buildHeap(arr)

	fmt.Println(arr)
	//堆排序
	for i := len(arr) -1 ; i > 0; i-- {
		//将堆顶与堆尾互换
		arr[0], arr[i] = arr[i], arr[0]
		//调整堆
		down(arr,0, i)
	}

	fmt.Println(arr)
}

func buildHeap(arr []int) {
	for i := len(arr)/2-1; i>=0 ; i-- {
		down(arr, i, len(arr))
	}
}

func down(arr []int, index, end int) {
	parent := index

	for {
		left := 2*parent+1
		if left >= end || left < 0 {
			break
		}
		min := left
		if right := left + 1; right < end && arr[right] < arr[left] {
			min = right
		}
		//小顶堆
		if arr[parent] < arr[min] {
			break
		}
		arr[parent], arr[min] = arr[min], arr[parent]

		parent = min
	}
}
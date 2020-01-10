//+build ignore
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
		child := 2*parent+1
		if child >= end || child < 0 {
			break
		}
		if child+1 < end && arr[child+1] < arr[child] {
			child++
		}
		//小顶堆
		if arr[parent] < arr[child] {
			break
		}
		arr[parent], arr[child] = arr[child], arr[parent]

		parent = child
	}
}
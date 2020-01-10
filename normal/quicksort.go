package normal

func QuickSort(arr []int, start, end int) {

	if start < end {
		mid := partion(arr, start, end)

		QuickSort(arr, start, mid-1)

		QuickSort(arr, mid+1, end)
	}

}

func partion(arr []int, i, j int) int {
	temp := arr[i]
	//todo 可以通过随机选择基准值进行排序  随机话快排
	//todo 可以通过随机在开头中部结尾三个值的中值作为基准  平衡快排
	//todo 可以先对有前部序部分进行跳过
	for {
		for ; i < j && arr[j] >= temp; j-- {
		}

		arr[i] = arr[j]

		for ; i < j && arr[i] <= temp; i++ {
		}

		arr[j] = arr[i]

		if i >= j {
			break
		}
	}

	arr[j] = temp

	return j
}

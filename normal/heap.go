package normal

type Student struct {
	name  string
	score int
}

type StudentHeap []Student

func (h StudentHeap) Len() int { return len(h) }

func (h StudentHeap) Less(i, j int) bool {
	return h[i].score < h[j].score //最小堆
	//return stu[i].score > stu[j].score //最大堆
}

func (h StudentHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *StudentHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Student))
}

func (h *StudentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

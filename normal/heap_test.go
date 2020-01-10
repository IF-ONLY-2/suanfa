package normal

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestUserHeap(t *testing.T) {
	ud := UserSlice{User{"a", 5}, User{"b", 2}, User{"c", 3}}

	heap.Init(&ud)

	fmt.Printf("%+v", ud)

	for len(ud) > 0 {
		fmt.Printf("\n%+v\n", heap.Pop(&ud))
	}

	fmt.Printf("%+v", ud)
}

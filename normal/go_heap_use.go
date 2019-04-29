package main

import (
	"container/heap"
	"fmt"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

func (US UserSlice)Len() int {return len(US)}

func (US UserSlice)Swap(i, j int) {US[i], US[j] = US[j], US[i]}

func (US UserSlice)Less(i, j int) bool {return US[i].Age < US[j].Age}

func (US *UserSlice)Push(x interface{})  {
	item := x.(User)
	*US = append(*US, item)
}

func (US *UserSlice)Pop() interface{} {
	item := (*US)[len(*US)-1]

	*US = (*US)[:len(*US)-1]

	return item
}

func main() {

	ud := UserSlice{User{"a",5}, User{"b",2},User{"c",3}}

	heap.Init(&ud)


	fmt.Printf("%+v",ud)

	for len(ud)>0 {
		fmt.Printf("\n%+v\n",heap.Pop(&ud))
	}

	fmt.Printf("%+v",ud)

}


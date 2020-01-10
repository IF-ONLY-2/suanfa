package normal

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (us UserSlice) Len() int { return len(us) }

func (us UserSlice) Swap(i, j int) { us[i], us[j] = us[j], us[i] }

func (us UserSlice) Less(i, j int) bool { return us[i].Age < us[j].Age }

func (us *UserSlice) Push(x interface{}) {
	item := x.(User)
	*us = append(*us, item)
}

func (us *UserSlice) Pop() interface{} {
	item := (*us)[len(*us)-1]

	*us = (*us)[:len(*us)-1]

	return item
}

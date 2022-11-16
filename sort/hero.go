package sort

import "sort"

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

// Less 决定使用什么标准来排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Sort() {
	sort.Sort(hs)
}

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less 决定使用什么标准来排序
func (s Students) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s Students) Sort() {
	sort.Sort(s)
}

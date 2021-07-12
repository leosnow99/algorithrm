package util

type SortInt []int

func (s SortInt) Len() int {
	return len(s)
}

func (s SortInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
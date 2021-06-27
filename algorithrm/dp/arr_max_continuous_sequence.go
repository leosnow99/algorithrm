package dp

import (
	util "algorithm/util"
)

/*
「数组中的最长连续序列」

给定无序数组arr，返回其中最长的连续序列的长度。
*/
func longestConsecutive(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := 1
	maps := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if _, ok := maps[arr[i]]; !ok {
			maps[arr[i]] = 1
			if _, ok = maps[arr[i]-1]; ok {
				max = util.Max(max, merge(maps, arr[i]-1, arr[i]))
			}
			if _, ok = maps[arr[i]+1]; ok {
				max = util.Max(max, merge(maps, arr[i], arr[i]+1))
			}
		}
	}
	return max
}

func merge(maps map[int]int, less, more int) int {
	left := less - maps[less] + 1
	right := more + maps[more] - 1
	length := right - left + 1
	maps[left] = length
	maps[right] = length
	return length
}

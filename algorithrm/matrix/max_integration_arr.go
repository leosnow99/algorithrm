package matrix

import (
	"algorithm/util"
	"math"
	"sort"
)

/**
最长的可整合子数组的长度
【题目】
先给出可整合数组的定义：如果一个数组在排序之后，每相邻两个数差的绝对值都为1，
则该数组为可整合数组。例如，[5,3,4,6,2]排序之后为[2,3,4,5,6]，符合每相邻两个数差的绝对值
都为1，所以这个数组为可整合数组。
给定一个整型数组arr，请返回其中最大可整合子数组的长度。例如，[5,5,3,2,6,4,3]的最大
可整合子数组为[5,3,2,6,4]，所以返回5。
*/

func getLIL1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	length := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if isIntegration(arr, i, j) {
				length = util.Max(length, j-i+1)
			}
		}
	}
	return length
}

func isIntegration(arr []int, left, right int) bool {
	newArr := arr[left : right+1]
	sort.Sort(util.SortInt(newArr))
	for i := 1; i < len(arr); i++ {
		if arr[i-1]+1 != arr[i] {
			return false
		}
	}
	return true
}

func getLIL2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	length, max, min := 0, 0, 0
	sets := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		max = math.MinInt32
		min = math.MaxInt32
		for j := i; j < len(arr); j++ {
			if _, ok := sets[arr[j]]; ok {
				break
			}
			sets[arr[j]] = true
			max = util.Max(max, arr[j])
			min = util.Min(min, arr[j])
			if max-min == j-i {
				length = util.Max(length, j-i+1)
			}
		}

		sets = make(map[int]bool)
	}
	return length
}

package matrix

import "algorithm/util"

/**
未排序数组中累加和为给定值的最长子数组系列问题
【题目】
给定一个无序数组arr，其中元素可正、可负、可0。给定一个整数k，求arr 所有的子数组中累加和为k 的最长子数组长度。

补充问题1：给定一个无序数组arr，其中元素可正、可负、可0。求arr 所有的子数组中正数与负数个数相等的最长子数组长度。
	第一个补充问题是先把数组arr 中的正数全部变成1，负数全部变成-1，0 不变，然后求累加和为0 的最长子数组长度即可。

补充问题2：给定一个无序数组arr，其中元素只是1 或0。求arr 所有的子数组中0 和1个数相等的最长子数组长度。
	先把数组arr 中的0 全部变成-1，1 不变，然后求累加和为0 的最长子数组长度即可
*/

func maxLength(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	maps := make(map[int]int)
	maps[0] = -1
	length, sum := 0, 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if step, ok := maps[sum-target]; ok {
			length = util.Max(length, i-step)
		}
		if _, ok := maps[sum]; !ok {
			maps[sum] = i
		}
	}

	return length
}

func maxLengthEqualPositiveAndNegative(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	for idx, number := range arr {
		if number > 0 {
			arr[idx] = 1
		} else if number < 0 {
			arr[idx] = - 1
		}
	}

	return maxLength(arr, 0)
}

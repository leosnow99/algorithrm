package matrix

import "algorithm/util"

/**
未排序正数数组中累加和为给定值的最长子数组长度
【题目】
给定一个数组arr，该数组无序，但每个值均为正数，再给定一个正数k。求arr 的所有子
数组中所有元素相加和为k 的最长子数组长度。
*/

func getMaxLength(arr []int, k int) int {
	if len(arr) == 0 || k <= 0 {
		return 0
	}

	left, right, sum, length := 0, 0, arr[0], 0
	for right < len(arr) {
		if sum == k {
			length = util.Max(length, right-left)
		} else if sum < k {
			right++
			if right == sum {
				break
			}
			sum += arr[right]
		} else {
			sum -= arr[left]
			left++
		}
	}
	return length
}

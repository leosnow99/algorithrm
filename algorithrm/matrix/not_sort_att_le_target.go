package matrix

import "algorithm/util"

/**
未排序数组中累加和小于或等于给定值的最长子数组长度
【题目】
给定一个无序数组arr，其中元素可正、可负、可0。给定一个整数k，求arr 所有的子数组中累加和小于或等于k 的最长子数组长度。
例如：arr=[3,-2,-4,0,6]，k=-2，相加和小于或等于-2 的最长子数组为{3,-2,-4,0}，所以结果返回4。
*/

func maxLeLength(arr []int, target int) int {
	h := make([]int, len(arr)+1)
	sum := 0
	h[0] = sum

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		h[i+1] = util.Max(sum, h[i])
	}

	sum, res, pre, length := 0, 0, 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		pre = getLessIndex(arr, sum-target)

		length = i - pre + 1
		if pre == - 1 {
			length = 0
		}
		res = util.Max(res, length)
	}

	return rs
}

// 二分查找大于或等于某一个值的累加和最早出现的位置
func getLessIndex(arr []int, num int) int {
	low, high, mid, res := 0, len(arr)-1, 0, -1
	for low <= high {
		mid = low + (high-low)>>2
		if arr[mid] >= num {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return res
}

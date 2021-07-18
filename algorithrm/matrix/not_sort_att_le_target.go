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

	return res
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

func maxLengthAwesome(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	minSums := make([]int, len(arr))
	minSumEnds := make([]int, len(arr))
	minSums[len(arr)-1] = arr[len(arr)-1]
	minSumEnds[len(arr)-1] = len(arr) - 1

	for i := len(arr) - 2; i != -1; i-- {
		if minSums[i+1] < 0 {
			minSums[i] = arr[i] + minSums[i+1]
			minSumEnds[i] = minSumEnds[i+1]
		} else {
			minSums[i] = arr[i]
			minSumEnds[i] = i
		}
	}

	end, sum, res := 0, 0, 0
	// i 是窗口最左的位置， end 是窗口最右的位置
	for i := 0; i < len(arr); i++ {
		// for 循环结束之后：
		// 1) 如果以i 开头的情况下，累加和小于或等于k 的最长子数组是arr[i..end-1]，看看这个子数组长度能不能更新res
		// 2) 如果以i 开头的情况下，累加和小于或等于k 的最长子数组比arr[i..end-1]短, 不管是否更新res，都不会影响最终结果
		for end < len(arr) && sum+minSums[end] <= k {
			sum += minSums[end]
			end = minSumEnds[end] + 1
		}

		res = util.Max(res, end-i)

		if end > i {
			sum -= arr[i]
		} else {
			end = i + 1
		}
	}

	return res
}

package matrix

/**
需要排序的最短子数组长度
【题目】
给定一个无序数组arr，求出需要排序的最短子数组长度。
例如：arr = [1,5,3,4,2,6,7]返回4，因为只有[5,3,4,2]需要排序。
*/

func getMinLength(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	min := arr[len(arr)-1]
	noMinIndex := -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > min {
			noMinIndex = i
		} else {
			min = arr[i]
		}
	}
	if noMinIndex == -1 {
		return 0
	}

	max := arr[0]
	noMaxIndex := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < max {
			noMaxIndex = arr[i]
		} else {
			max = arr[i]
		}
	}
	return noMaxIndex - noMinIndex
}

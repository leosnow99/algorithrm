package matrix

import (
	"algorithm/util"
	"math"
)

/**
子数组的最大累加和问题
【题目】
	给定一个数组arr，返回子数组的最大累加和。
	例如，arr=[1,-2,3,5,-2,6,-1]，所有的子数组中，[3,5,-2,6]可以累加出最大的和12，所以返回
	12。
*/

func maxSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := math.MinInt32
	cur := 0
	for i := 0; i < len(arr); i++ {
		cur += arr[i]
		max = util.Max(max, cur)
		if cur < 0 {
			cur = 0
		}
	}

	return max
}

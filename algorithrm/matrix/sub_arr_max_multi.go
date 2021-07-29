package matrix

import "algorithm/util"

/**
数组中子数组的最大累乘积

【题目】
给定一个double 类型的数组arr，其中的元素可正、可负、可0，返回子数组累乘的最大乘积。
例如，arr=[-2.5，4，0，3，0.5，8，-1]，子数组[3，0.5，8]累乘可以获得最大的乘积12，所以返回12。
*/

func maxProduct(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max, min, res := arr[0], arr[0], 0
	maxEnd, minEnd := 0, 0

	for i := 1; i < len(arr); i++ {
		maxEnd = max * arr[i]
		max = util.Max(max, maxEnd)

		minEnd = min * arr[i]
		min = util.Min(min, minEnd)

		res = util.Max(res, max)
	}

	return res
}

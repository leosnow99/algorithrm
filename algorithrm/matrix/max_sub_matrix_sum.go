package matrix

import (
	"algorithm/util"
	"math"
)

/**
子矩阵的最大累加和问题
【题目】
给定一个矩阵matrix，其中的值有正、有负、有0，返回子矩阵的最大累加和。
*/

// 也就是说，如果一个矩阵一共有k 行且限定必须含有k 行元素的情况下，我们只要把矩阵中每一列的k 个元素累加生成一个累加数组，
// 然后求出这个数组的最大累加和，这个最大累加和就是必须含有k 行元素的子矩阵中的最大累加和。

func maxSubMatrixSum(m [][]int) int {
	if len(m) == 0 || len(m[0]) == 0 {
		return 0
	}

	max, cur := math.MinInt32, 0
	s := make([]int, len(m[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			cur = 0
			for k := 0; k < len(m[0]); k++ {
				s[k] += m[j][k]
				cur += s[k]
				max = util.Max(max, cur)
				if cur < 0 {
					return cur
				}
			}
		}
	}

	return max
}

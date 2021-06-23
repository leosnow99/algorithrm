package string_type

import (
	"algorithm/src/util"
	"math"
)

/**
回文最少分割数

给定一个字符串str，返回把str 全部切成回文子串的最小分割数。
*/

func minCut(str string) int {
	if len(str) == 0 {
		return 0
	}
	chas := []byte(str)
	lens := len(chas)
	dp := make([]int, lens+1)
	dp[lens] = -1
	p := make([][]bool, lens)
	for i := range p {
		p[i] = make([]bool, lens)
	}
	for i := lens - 1; i >= 0; i-- {
		dp[i] = math.MaxInt32
		for j := i; j < lens; j++ {
			if (chas[i] == chas[j]) && (j-i < 2 || p[i+1][j-1]) {
				p[i][j] = true
				dp[i] = util.Min(dp[i], dp[j+1]+1)
			}
		}
	}
	return dp[0]
}

package dp

import "fmt"

/**
给定一个数组arr，代表一排有分数的气球。每打爆一个气球都能获得分数，假设打爆气球
的分数为X，获得分数的规则如下：

1）如果被打爆气球的左边有没被打爆的气球，找到离被打爆气球最近的气球，假设分数为
L；如果被打爆气球的右边有没被打爆的气球，找到离被打爆气球最近的气球，假设分数为R。
获得分数为L×X×R。

2）如果被打爆气球的左边有没被打爆的气球，找到离被打爆气球最近的气球，假设分数为
L；如果被打爆气球的右边所有气球都已经被打爆。获得分数为L×X。

3）如果被打爆气球的左边所有的气球都已经被打爆；如果被打爆气球的右边有没被打爆的
气球，找到离被打爆气球最近的气球，假设分数为R；如果被打爆气球的右边所有气球都已经
被打爆。获得分数为X×R。

4）如果被打爆气球的左边和右边所有的气球都已经被打爆。获得分数为X。
目标是打爆所有气球，获得每次打爆的分数。通过选择打爆气球的顺序，可以得到不同的
总分，请返回能获得的最大分数
*/
func process(arr []int, L, R int) int {
	if L == R {
		return arr[L-1] * arr[L] * arr[R+1]
	}

	// 最后打爆左右边界的问题
	maxLeft := (arr[L-1] * arr[L] * arr[R+1]) + process(arr, L+1, R)
	maxRight := (arr[L-1] * arr[R] * arr[R+1]) + process(arr, L, R-1)
	max := maxLeft
	if maxLeft < maxRight {
		max = maxRight
	}

	// 中间位置气球最后被打爆的每一种方案
	for i := L + 1; i < R; i++ {
		tmp := (arr[L-1] * arr[i] * arr[R+1]) + process(arr, L, i-1) + process(arr, i+1, R)
		if max < tmp {
			max = tmp
		}
	}

	return max
}

func maxCoins(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}

	var help []int
	help = append(help, 1)
	help = append(help, arr...)
	help = append(help, 1)
	return process(help, 1, len(arr))
}

func MaxCoins2(arr []int) int {
	N := len(arr)
	if N == 0 {
		return 0
	}
	if N == 1 {
		return arr[0]
	}

	var help []int
	help = append(help, 1)
	help = append(help, arr...)
	help = append(help, 1)
	var dp = make([][]int, N+2)
	for row := range dp {
		dp[row] = make([]int, N+2)
	}
	for i := 1; i <= N; i++ {
		dp[i][i] = help[i-1] * help[i] * help[i+1]
	}

	for L := N; L >= 1; L-- {
		for R := L + 1; R <= N; R++ {
			// 求解dp[L][R], 表示help[L...R]上打爆所有气球的最大分数
			// 最后打爆help[L] 的方案
			finalL := (help[L-1] * help[L] * help[R+1]) + dp[L+1][R]
			// 最后打爆help[R] 的方案
			finalR := (help[L-1] * help[R] * help[R+1]) + dp[L][R-1]
			dp[L][R] = finalL
			if finalL < finalR {
				dp[L][R] = finalR
			}
			// 打爆中间值的每一种方案
			for i := L + 1; i < R; i++ {
				tmp := (help[L-1] * help[i] * help[R+1]) + dp[L][i-1] + dp[i+1][R]
				if dp[L][R] < tmp {
					dp[L][R] = tmp
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[1][N]
}

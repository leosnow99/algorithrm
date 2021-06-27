package dp

/**
给定数组 arr，arr 中所有的值都为正数且不重复。每个值代表一种面值的货币，每种面值
的货币可以使用任意张，再给定一个整数aim，代表要找的钱数，求换钱有多少种方法。
*/

// 暴力递归方法
func coins1(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	return process1(arr, 0, aim)
}

func process1(arr []int, idx, aim int) int {
	res := 0
	if idx == len(arr) {
		if aim == 0 {
			res = 1
		}
		res = 0
	} else {
		for i := 0; arr[idx]*i <= aim; i++ {
			res += process1(arr, idx+1, aim-arr[idx]*i)
		}
	}
	return res
}

// 记忆化搜索
// 记忆化搜索的方法是针对暴力递归最初级的优化技巧，分析递归函数的状态可以由哪些变量表示，做出相应维度和大小的map 即可
func coins2(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}
	maps := make([][]int, len(arr)+1)
	for idx := range maps {
		maps[idx] = make([]int, aim+1)
	}
	return process1(arr, 0, aim)
}

func process2(arr []int, idx, aim int, maps [][]int) int {
	res := 0
	if idx == len(arr) {
		if aim == 0 {
			res = 1
		}
	} else {
		mapValue := 0
		for i := 0; arr[idx]*i <= aim; i++ {
			mapValue = maps[i+1][aim-arr[idx]*i]
			if mapValue != 0 {
				if mapValue == -1 {
					res = 0
				}
				res = mapValue
			} else {
				res += process2(arr, idx+1, aim-arr[idx]*i, maps)
			}
		}
	}
	if res == 0 {
		maps[idx][aim] = -1
	} else {
		maps[idx][aim] = res
	}
	return res
}

// 动态规划
func coins3(arr []int, aim int) int {
	if len(arr) == 0 || aim < 0 {
		return 0
	}

	dp := make([][]int, len(arr))
	for idx := range dp {
		dp[idx] = make([]int, aim+1)
	}
	for i := 0; i < len(arr); i++ {
		dp[i][0] = 1
	}
	for i := 1; arr[0]*i <= aim; i++ {
		dp[0][arr[0]*i] = 1
	}

	num := 0
	for i := 1; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			num = 0
			for k := 0; j-arr[i]*k >= 0; k++ {
				num += dp[i-1][j-arr[i]*k]
			}
			dp[i][j] = num
		}
	}
	return dp[len(arr)-1][aim]
}

// 动态规划 优化版本
func coins4(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}

	dp := make([][]int, len(arr))
	for idx := range dp {
		dp[idx] = make([]int, aim+1)
	}
	for i := 0; i < len(arr); i++ {
		dp[i][0] = 1
	}
	for i := 1; arr[0]*i <= aim; i++ {
		dp[0][arr[0]*i] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			dp[i][j] = dp[i-1][j]
			if j-dp[i][j] > 0 {
				dp[i][j] += dp[i][j-arr[i]]
			}
		}
	}
	return dp[len(arr)-1][aim]
}

// 在优化版本之上空间压缩
func coins5(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}

	dp := make([]int, aim+1)
	for i := 0; arr[0]*i <= aim; i++ {
		dp[i] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			if j-arr[j] >= 0 {
				dp[j] += dp[j-arr[i]]
			}
		}
	}

	return dp[aim]
}

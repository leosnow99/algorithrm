package dp

//数组异或和的定义：把数组中所有的数异或起来得到的值。
//给定一个整型数组arr，其中可能有正、有负、有零。你可以随意把整个数组切成若干个不
//相容的子数组，求异或和为0 的子数组最多能有多少个？
func mostEOR(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	eor := 0
	maps := make(map[int]int)
	maps[0] = -1
	dp := make([]int, len(arr))
	if arr[0] == 0 {
		dp[0] = 1
	} else {
		dp[0] = 0
	}
	maps[arr[0]] = 0
	for i := 1; i < len(arr); i++ {
		eor ^= arr[i]
		if preEorIndex, ok := maps[eor]; ok {
			if preEorIndex == -1 {
				dp[i] = 1
			} else {
				dp[i] = dp[preEorIndex] + 1
			}
		}
		if dp[i] < dp[i-1] {
			dp[i] = dp[i-1]
		}
		maps[eor] = i
	}
	return dp[len(dp)-1]
}

package dynamic_programing

//爬楼梯
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//爬楼梯2
func climbStairsDemo2(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	first := 1
	last := 2
	for i := 3; i < n+1; i++ {
		first, last = last, first+last
	}
	return last
}

//最大子序和
func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > result {
			result = nums[i]
		}
	}
	return result
}

//最长上升子序列
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return 1
	}

	dp := make([]int, length)
	result := 1
	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//三角形最小路径和
func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return triangle[0][0]
	}

	result := 2<<31 - 1

	triangle[1][0] += triangle[0][0]
	triangle[1][1] += triangle[0][0]

	for i := 2; i < length; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == (len(triangle[i]) - 1) {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] = min(triangle[i-1][j], triangle[i-1][j+1]) + triangle[i][j]
			}
		}
	}

	for _, value := range triangle[length-1] {
		result = min(result, value)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//最小路径和
func minPathSum(grid [][]int) int {
	length := len(grid)
	if length < 1 {
		return 0
	}
	for i := 0; i < length; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			} else if j == 0 && i != 0 {
				grid[i][j] += grid[i-1][j]
			} else {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[length-1][len(grid[0])-1]
}

//打家劫舍
func rob(nums []int) int {
	size := len(nums)
	if size < 1 {
		return 0
	} else if size == 1 {
		return nums[0]
	} else if size == 2 {
		return max(nums[0], nums[1])
	}

	nums[1] = max(nums[0], nums[1])
	for i := 2; i < size; i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[size-1]

}

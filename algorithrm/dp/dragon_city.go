package dp

import (
	"algorithm/util"
)

// 龙与地下城游戏问题
// 给定一个二维数组map，含义是一张地图
// 游戏的规则如下：
//	 骑士从左上角出发，每次只能向右或向下走，最后到达右下角见到公主。
//	 地图中每个位置的值代表骑士要遭遇的事情。如果是负数，说明此处有怪兽，要让骑士损失血量。如果是非负数，代表此处有血瓶，能让骑士回血。
//	 骑士从左上角到右下角的过程中，走到任何一个位置时，血量都不能少于1。
// 为了保证骑士能见到公主，初始血量至少是多少？根据map，返回初始血量。
func minHP1(arr [][]int) int {
	if arr == nil || arr[0] == nil {
		return 1
	}
	// 初始化 dp 数组
	row := len(arr)
	col := len(arr[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	row--
	col--
	dp[row][col] = arr[row][col]
	if arr[row][col] < 0 {
		dp[row][col] = -arr[row][col] + 1
	}
	// 处理最后一行情况
	for i := col - 1; i >= 0; i-- {
		dp[row][i] = 1
		if dp[row][i] < dp[row][i+1]-arr[row][i] {
			dp[row][i] = dp[row][i+1] - arr[row][i]
		}
	}

	right, down := 0, 0
	for i := row - 1; i >= 0; i-- {
		dp[i][col] = util.Max(dp[i+1][col]-dp[i][col], 1)
		for j := col - 1; j >= 0; j-- {
			right = util.Max(dp[i][j+1]-arr[i][j], 1)
			down = util.Max(dp[i+1][j]-arr[i][j], 1)
			dp[i][j] = util.Min(right, down)
		}

	}

	return dp[0][0]
}

func minHP2(arr [][]int) int {
	if arr == nil || arr[0] == nil {
		return 1
	}
	// 初始化 dp 数组
	more := len(arr)
	less := len(arr[0])
	if more < less {
		more, less = less, more
	}
	rowMore := more == len(arr)
	dp := make([]int, less)
	tmp := arr[len(arr)-1][len(arr[0])-1]
	if tmp > 0 {
		tmp = 1
	}
	dp[less-1] = -tmp + 1
	row, col := 0, 0
	for i := less - 2; i >= 0; i-- {
		row = i
		col = more - 1
		if rowMore {
			row, col = col, row
		}
		dp[i] = util.Max(dp[i+1]-arr[row][col], 1)
	}
	chosen1, chosen2 := 0, 0
	for i := more - 2; i >= 0; i-- {
		row = less - 1
		col = i
		if rowMore {
			row, col = col, row
		}
		dp[less-1] = util.Max(dp[less-1]-arr[row][col], 1)
		for j := less - 2; j >= 0; j-- {
			row = j
			col = i
			if rowMore {
				row, col = col, row
			}
			chosen1 = util.Max(dp[j]-arr[row][col], 1)
			chosen2 = util.Max(dp[j+1]-arr[row][col], 1)
			dp[j] = util.Min(chosen1, chosen2)
		}
	}
	return dp[0]
}

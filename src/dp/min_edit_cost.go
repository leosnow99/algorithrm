package dp

//给定两个字符串str1 和str2，再给定三个整数ic、dc 和rc，分别代表插入、删除和替换一
//个字符的代价，返回将str1 编辑成str2 的最小代价。
func minCost1(str1, str2 string, ic, dc, rc int) int {
	chs1 := []byte(str1)
	chs2 := []byte(str2)
	row := len(chs1) + 1
	col := len(chs2) + 1
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		dp[i][0] = dc * i
	}
	for i := 1; i < col; i++ {
		dp[0][i] = ic * i
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			// 替换
			if chs1[i-1] == chs2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rc
			}
			// 添加
			if (dp[i][j-1] + ic) < dp[i][j] {
				dp[i][j] = dp[i][j-1] + ic
			}
			// 删除
			if (dp[i-1][j] + dc) < dp[i][j] {
				dp[i][j] = dp[i-1][j] + dc
			}
		}
	}
	return dp[row-1][col-1]
}

// 空间压缩
func minCost2(str1, str2 string, ic, rc, dc int) int {
	longs := []byte(str1)
	shorts := []byte(str2)
	if len(longs) < len(shorts) {
		longs, shorts = shorts, longs
	}
	dp := make([]int, len(shorts)+1)
	for i := 1; i <= len(longs); i++ {
		dp[i] = ic * i
	}
	for i := 1; i <= len(longs); i++ {
		pre := dp[0]
		dp[0] = dc * i
		for j := 1; j <= len(shorts); j++ {
			tmp := dp[j]
			if longs[i-1] == shorts[j-1] {
				dp[j] = pre
			} else {
				dp[j] = pre + rc
			}
			if dp[j-1]+ic < dp[j] {
				dp[j] = dp[j-1] + ic
			}
			if tmp+dc < dp[j] {
				dp[j] = tmp + dc
			}
			pre = tmp
		}
	}
	return dp[len(dp)-1]
}

package dp

import "fmt"

/**
给定两个字符串str1 和str2，返回两个字符串的最长公共子序列。
str1="1A2C3D4B56"，str2="B1D23CA45B6A"。
"123456"或者"12C4B6"都是最长公共子序列，返回哪一个都行。
*/

/**
如果str1 的长度为M，str2 的长度为N，生成大小为M×N 的矩阵dp，行数为M，列数为N。
dp[i][j]的含义是str1[0..i]与str2[0..j]的最长公共子序列的长度。
从左到右，再从上到下计算矩阵dp。
*/
func getdp(str1, str2 []byte) [][]int {
	dp := make([][]int, len(str1))
	for i := range dp {
		dp[i] = make([]int, len(str2))
	}

	if str1[0] == str2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < len(str1); i++ {
		dp[i][0] = dp[i-1][0]
		if str1[i] == str2[0] {
			dp[i][0] = 1
		}
	}
	for i := 1; i < len(str2); i++ {
		dp[0][i] = dp[0][i-1]
		if str2[i] == str1[0] {
			dp[0][i] = 1
		}
	}
	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			dp[i][j] = dp[i-1][j]
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i][j-1]
			}
			if str1[i] == str2[j] && dp[i][j] < (dp[i-1][j-1]+1) {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return dp
}

func Lcse(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	chs1 := []byte(str1)
	chs2 := []byte(str2)
	dp := getdp(chs1, chs2)
	m, n := len(chs1)-1, len(chs2)-1
	res := make([]byte, dp[m][n])
	fmt.Println(dp)
	for index := len(res) - 1; index >= 0; {
		if n > 0 && dp[m][n] == dp[m][n-1] {
			n--
		} else if m > 0 && dp[m][n] == dp[m-1][n] {
			m--
		} else {
			res[index] = chs1[m]
			index--
			m--
			n--
		}
	}
	return string(res)
}

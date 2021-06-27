package dp

/**
给定两个字符串str1 和str2，返回两个字符串的最长公共子串。
str1="1AB2345CD"，str2="12345EF"，返回"2345"。
*/

func getMCSDP(str1, str2 []byte) [][]int {
	dp := make([][]int, len(str1))
	for i := range dp {
		dp[i] = make([]int, len(str2))
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1
		}
	}
	for i := 0; i < len(str2); i++ {
		if str2[i] == str1[0] {
			dp[0][i] = 1
		}
	}
	for i := 1; i < len(str1); i++ {
		for j := 1; j < len(str2); j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}
	return dp
}

func Lcst1(str1, str2 string) string {
	if str1 == "" && str2 == "" {
		return ""
	}
	chs1, chs2 := []byte(str1), []byte(str2)
	dp := getMCSDP(chs1, chs2)
	end, max := 0, 0
	for i := 0; i < len(chs1); i++ {
		for j := 0; j < len(chs2); j++ {
			if max < dp[i][j] {
				max = dp[i][j]
				end = i
			}
		}
	}
	return string([]byte(str1)[end-max+1 : end+1])
}

func Lcst2(str1, str2 string) string {
	if str1 == "" && str2 == "" {
		return ""
	}
	chs1, chs2 := []byte(str1), []byte(str2)
	row := 0             // 斜线开始位置的行
	col := len(chs2) - 1 // 斜线开始位置的列
	max := 0             // 记录最大长度
	end := 0             // 最大长度更新时，记录子串的结束位置
	for row < len(chs1) {
		i, j, length := row, col, 0
		for i < len(chs1) && j < len(chs2) {
			if chs1[i] != chs2[j] {
				length = 0
			} else {
				length++
			}
			if max < length {
				end = i
				max = length
			}
			i++
			j++
		}
		if col > 0 { //斜线开始位置的先向左移动
			col--
		} else { // 列移动到最左之后，行向下移动
			row++
		}
	}
	return string([]byte(str1)[end-max+1 : end+1])
}

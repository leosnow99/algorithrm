package dp

//给定三个字符串str1、str2 和aim，如果aim 包含且仅包含来自str1 和str2 的所有字符，而
//且在aim 中属于str1 的字符之间保持原来在str1 中的顺序，属于str2 的字符之间保持原来在str2
//中的顺序，那么称aim 是str1 和str2 的交错组成。实现一个函数，判断aim 是否是str1 和str2
//交错组成。

func isCross(str1, str2, aim string) bool {
	ch1 := []byte(str1)
	ch2 := []byte(str2)
	chAim := []byte(aim)
	if len(ch1)+len(ch2) != len(aim) {
		return false
	}
	dp := make([][]bool, len(ch1)+1)
	for i := 0; i <= len(ch1); i++ {
		dp[i] = make([]bool, len(ch2)+1)
	}
	for i := 1; i <= len(ch1); i++ {
		if ch1[i-1] != chAim[i-1] {
			break
		}
		dp[i][0] = true
	}
	for i := 1; i <= len(ch2); i++ {
		if ch2[i-1] != chAim[i-1] {
			break
		}
		dp[0][i-1] = true
	}
	for i := 1; i <= len(ch1); i++ {
		for j := 1; j <= len(ch2); j++ {
			if ch1[i-1] == chAim[i+j-1] && dp[i-1][j] || ch2[j-1] == chAim[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}
	return dp[len(ch1)][len(ch2)]
}

func isCross2(str1, str2, aim string) bool {
	longs := []byte(str1)
	shorts := []byte(str2)
	chAim := []byte(aim)
	if len(longs)+len(shorts) != len(chAim) {
		return false
	}
	if len(longs) < len(shorts) {
		longs, shorts = shorts, longs
	}
	dp := make([]bool, len(shorts)+1)
	for i := 1; i <= len(shorts); i++ {
		if shorts[i-1] != chAim[i-1] {
			break
		}
		dp[i] = true
	}
	for i := 1; i < len(longs); i++ {
		dp[0] = dp[0] && longs[i-1] == chAim[i-1]
		for j := 1; j <= len(shorts); j++ {
			if longs[i-1] == chAim[i+j-1] && dp[j] || shorts[j-1] == chAim[i+j-1] && dp[j-1] {
				dp[j] = true
			} else {
				dp[j] = false
			}
		}
	}
	return dp[len(shorts)]
}

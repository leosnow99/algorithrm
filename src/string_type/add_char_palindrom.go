package string_type

import (
	"algorithm/src/util"
)

// 添加最少字符使字符串整体都是回文字符串
// 给定一个字符串str，如果可以在str 的任意位置添加字符，请返回在添加字符最少的情况下，让 str 整体都是回文字符串的一种结果。

// 如果可以在str 的任意位置添加字符，最少需要添几个字符可以让str 整体都是回文字符串。这个问题可以用动态规划的方法求解。
// 如果str 的长度为N，动态规划表是一个N×N 的矩阵，记为dp[][]。dp[i][j]值的含义代表子串str[i..j]最少添加几个字符可以使str[i..j]整体都是回文串。
func getDp(str []byte) [][]int {
	dp := make([][]int, len(str))
	for i := range dp {
		dp[i] = make([]int, len(str))
	}
	for j := 1; j < len(str); j++ {
		if str[j-1] != str[j] {
			dp[j-1][j] = 1
		}
		for i := j - 2; i > -1; i-- {
			if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = util.Min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp

}

// 在添加字符最少的情况下，让str 整体都是回文字符串的一种结果。
// 首先，dp[0][N-1]的值代表整个字符串最少需要添加几个字符，所以，如果最后的结果
// 记为字符串res，res 的长度=dp[0][N-1]+str 的长度，然后依次设置res 左右两头的字符。
func getPalindrome1(str string) string {
	if len(str) < 2 {
		return str
	}
	chas := []byte(str)
	dp := getDp(chas)
	i, j := 0, len(chas)-1
	resL, resR := 0, len(chas)

	var res = make([]byte, len(chas)+dp[0][len(chas)-1])
	for i <= j {
		if chas[i] == chas[j] {
			res[resL] = chas[i]
			resL++
			i++
			res[resR] = chas[j]
			resR--
			j--
		} else if dp[i][j-1] < dp[i+1][j] {
			res[resL] = chas[j]
			resL++
			res[resR] = chas[j]
			resR--
			j--
		} else {
			res[resL] = chas[i]
			resL++
			res[resR] = chas[i]
			resR--
			i++
		}
	}
	return string(res)
}

// 给定一个字符串str，再给定str的, 最长回文子序列 字符串strlps，请返回在添加
// 字符最少的情况下，让str整体都是回文字符串的一种结果。进阶问题比原问题多了一个参数， 请做到时间复杂度比原问题的实现低。
func getPalindrome2(str, strlps string) string {
	if len(str) == 0 {
		return str
	}
	chas := []byte(str)
	lps := []byte(strlps)
	var res = make([]byte, 2*len(str)-len(strlps))
	chasL, chasR := 0, len(chas)-1
	lpsL, lpsR := 0, len(lps)-1
	resL, resR := 0, len(res)
	tmpL, tmpR := 0, 0
	for lpsL < lpsR {
		tmpL, tmpR = chasL, chasR
		for chas[chasL] != lps[lpsL] {
			chasL++
		}
		for chas[chasR] != lps[lpsR] {
			chasR--
		}
		set(res, resL, resR, chas, chasL, tmpL, chasR, tmpR)
		resL += chasL - tmpL + tmpR - chasR
		resR -= chasL - tmpL + tmpR - chasR
		res[resL] = chas[chasL]
		res[resR] = chas[chasR]
		resL++
		resR--
		chasL++
		chasR--
		lpsL++
		lpsR--
	}
	return string(res)
}

func set(res []byte, resL, resR int, chas []byte, ls, le, rs, re int) {
	for i := ls; i < le; i++ {
		res[resL] = chas[i]
		res[resR] = chas[i]
		resL++
		resR--
	}
	for i := rs; i < re; i++ {
		res[resL] = chas[i]
		res[resR] = chas[i]
		resL++
		resR--
	}
}

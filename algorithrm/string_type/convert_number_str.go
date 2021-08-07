package string_type

import "math"

// 将整数字符串转成整数值
// 给定一个字符串str，如果str 符合日常书写的整数形式，并且属于32 位整数的范围，返回
// str 所代表的整数值，否则返回0。

// 1．如果str 不以“-”开头，也不以数字字符开头，例如，str=="A12"，返回false。
// 2．如果str 以“-”开头，但是str 的长度为1，即str=="-"，返回false。如果str 的长度大
//	 于1，但是“-”的后面紧跟着“0”，例如，str=="-0"或"-012"，返回false。
// 3．如果str 以“0”开头，但是str 的长度大于1，例如，str=="023"，返回false。
// 4．如果经过步骤1~步骤3 都没有返回，接下来检查str[1..N-1]是否都是数字字符，如果有
//	 一个不是数字字符，则返回false。如果都是数字字符，说明str 符合日常书写，返回true。
func isValid(chas []byte) bool {
	if chas[0] != '-' && chas[0] < '0' || chas[0] > '9' {
		return false
	}
	if chas[0] == '-' && (len(chas) == 0 || chas[1] == '0') {
		return false
	}
	if chas[0] == '0' && len(chas) > 1 {
		return false
	}
	for i := 1; i < len(chas); i++ {
		if chas[i] < '0' || chas[i] > '9' {
			return false
		}
	}
	return true
}

func convert(str string) int {
	if len(str) == 0 {
		return 0
	}
	chas := []byte(str)
	if !isValid(chas) {
		return 0
	}

	var posi bool
	if chas[0] != '-' {
		posi = true
	}
	minq := math.MinInt32 / 10
	minr := math.MinInt32 % 10
	var i = 0
	if !posi {
		i = 1
	}

	cur, res := byte(0), 0
	for ; i < len(chas); i++ {
		cur = '0' - chas[i]
		if res < minq || (res == minq && int(cur) < minr) {
			return 0
		}
		res = res*10 + int(cur)
	}

	if posi {
		return -res
	}

	return res
}

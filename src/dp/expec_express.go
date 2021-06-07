package dp

/**
表达式得到期望结果的组成种数

给定一个只由0（假）、1（真）、&（逻辑与）、|（逻辑或）和^（异或）五种字符组成的字
符串express，再给定一个布尔值desired。返回express 能有多少种组合方式，可以达到desired
的结果。
*/

/**
 表达式的长度必须是奇数。
 表达式下标为偶数位置的字符一定是'0'或者'1'。
 表达式下标为奇数位置的字符一定是'&'、'|'或'^'。
*/
func isValid(exp []byte) bool {
	if (len(exp) & 1) == 0 {
		return false
	}
	for i := 0; i < len(exp); i += 2 {
		if exp[i] != '0' && exp[i] != '1' {
			return false
		}
	}
	for i := 1; i < len(exp); i += 2 {
		if exp[i] != '&' && exp[i] != '^' && exp[i] != '|' {
			return false
		}
	}
	return true
}

func violenceDiscard(str string, desired bool) int {
	if str == "" {
		return 0
	}
	chs := []byte(str)
	if !isValid(chs) {
		return 0
	}
	return p(chs, desired, 0, len(chs))
}

func p(exp []byte, desired bool, l, r int) int {
	if l == r {
		if exp[l] == 1 {
			if desired {
				return 1
			} else {
				return 0
			}
		} else {
			if desired {
				return 0
			} else {
				return 1
			}
		}
	}
	res := 0
	if desired {
		for i := l + 1; i < r; i += 2 {
			switch exp[i] {
			case '&':
				res += p(exp, true, l, i-1) + p(exp, true, i+1, r)
				break
			case '^':
				res += p(exp, true, l, i-1) + p(exp, false, i+1, r)
				res += p(exp, false, l, i-1) + p(exp, true, i+1, r)
				break
			case '|':
				res += p(exp, true, l, i-1) + p(exp, false, i+1, r)
				res += p(exp, false, l, i-1) + p(exp, true, i+1, r)
				res += p(exp, true, l, i-1) + p(exp, true, i+1, r)
				break
			}
		}
	} else {
		for i := l + 1; i < r; i += 2 {
			switch exp[i] {
			case '&':
				res += p(exp, true, l, i-1) + p(exp, false, i+1, r)
				res += p(exp, false, l, i-1) + p(exp, true, i+1, r)
				res += p(exp, false, l, i-1) + p(exp, false, i+1, r)
				break
			case '^':
				res += p(exp, true, l, i-1) + p(exp, true, i+1, r)
				res += p(exp, false, l, i-1) + p(exp, false, i+1, r)
				break
			case '|':
				res += p(exp, true, l, i-1) + p(exp, true, i+1, r)
				break
			}
		}
	}
	return res
}

// 动态规划版本
func eeDP(exp string, desired bool) int {
	if exp == "" {
		return 0
	}
	chs := []byte(exp)
	if !isValid(chs) {
		return 0
	}
	t := make([][]int, len(chs))
	f := make([][]int, len(chs))
	for i := 0; i < len(chs); i++ {
		t[i] = make([]int, len(chs))
		f[i] = make([]int, len(chs))
	}
	if exp[0] == '1' {
		t[0][0] = 1
		f[0][0] = 0
	} else {
		t[0][0] = 0
		f[0][0] = 1
	}
	for i := 2; i < len(chs); i += 2 {
		if exp[i] == 1 {
			t[i][i] = 1
			f[i][i] = 0
		} else {
			t[i][i] = 0
			f[i][i] = 1
		}
		for j := i - 2; j >= 0; j -= 2 {
			for k := j; k < i; k += 2 {
				if chs[k+1] == '&' {
					t[j][i] += t[j][k] * t[k+2][i]
					f[j][i] += (f[j][k]+t[j][k])*f[k+2][i] + f[j][k]*t[k+2][i]
				} else if chs[k+1] == '|' {
					t[j][i] += (t[j][k]+f[j][k])*t[k+2][i] + t[j][k]*f[k+2][i]
					f[j][i] += f[j][k] * f[k+2][i]
				} else {
					t[j][i] += t[j][k]*f[k+2][i] + f[j][k]*t[k+2][i]
					f[j][i] += f[j][k]*f[k+2][i] + t[j][k]*t[k+2][i]
				}
			}
		}
	}
	if desired {
		return t[0][len(chs) -1]
	}
	return f[0][len(chs) -1]
}

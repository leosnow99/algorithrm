package string_type

/**
字符串匹配问题
给定字符串str，其中绝对不含有字符'.'和'*'。再给定字符串exp，其中可以含有'.'或'*'，'*'
字符不能是exp 的首字符，并且任意两个'*'字符不相邻。exp 中的'.'代表任何一个字符，exp 中
的'*'表示'*'的前一个字符可以有0 个或者多个。请写一个函数，判断str 是否能被exp 匹配。
*/

func strIsValid(s, e []byte) bool {
	for _, c := range s {
		if c == '*' || c == '.' {
			return false
		}
	}
	for i, c := range e {
		if c == '*' && (i == 0 || e[i-1] == '*') {
			return false
		}
	}
	return true
}

// 递归解法
func strIsMatch(str, exp string) bool {
	if len(str) == 0 || len(exp) == 0 {
		return false
	}
	s := []byte(str)
	e := []byte(exp)
	return processMatch(s, e, 0, 0)
}

// 从str 的si 位置开始，一直到str 结束位置的子串，即str[si...len(s)]，
// 是否能被从exp 的ei 位置开始一直到exp 结束位置的子串（即exp[ei...elen]）匹配
func processMatch(s, e []byte, si, ei int) bool {
	// 如果ei 为exp 的结束位置（ei==elen），si 也是str 的结束位置，返回true，因为“”可以匹配“”
	if ei == len(e) {
		return si == len(s)
	}

	// 如果ei 位置的下一个字符（e[ei+1]）不为'*'。那么就必须关注str[si]字符能否和exp[ei]字符匹配。
	// 如果str[si]与exp[ei]能匹配（e[ei] == s[si] || e[ei] == '.'），还要关注str 后续的部分能
	//  否被exp 后续的部分匹配，即process(s,e,si+1,ei+1)的返回值。
	if ei+1 == len(e) || e[ei+1] != '*' {
		return si != len(s) && (e[ei] == s[si] || e[ei] == '.') && processMatch(s, e, si+1, ei+1)
	}

	// 如果当前ei 位置的下一个字符（e[ei+1]）为'*'字符。
	for si != len(s) && (e[ei] == s[si] || e[ei] == '.') {
		if processMatch(s, e, si, ei+2) {
			return true
		}
		si++
	}

	return processMatch(s, e, si, ei+2)
}

// 因为dp[i][j]只依赖dp[i+1][j+1]或者dp[i+k][j+2](k≥0)的值，所以在单独计算完最后一行、最后一列与倒数第二列之后，
// 剩下的位置在从右到左，再从下到上计算dp 值的时候，所有依赖的值都被计算出来，直接拿过来用即可。
// 如果str 的长度为N，exp 的长度为M，因为有枚举的过程，所以时间复杂度为O(N2×M)，额外空间复杂度为O(N×M)。
func isMatchDp(str, exp string) bool {
	if len(str) == 0 || len(exp) == 0 {
		return false
	}
	s := []byte(str)
	e := []byte(exp)
	if !strIsValid(s, e) {
		return false
	}

	dp := initDpMap(s, e)
	for i := len(s) - 1; i > -1; i-- {
		for j := len(e) - 2; i > -1; i-- {
			if e[j+1] != '*' {
				dp[i][j] = (s[i] == e[j] || e[j] == '.') && dp[i+1][j+1]
			} else {
				si := i
				for si != len(s) && (s[si] == e[j] || e[j] == '.') {
					if dp[si][j+2] {
						dp[i][j] = true
						break
					}
					si++
				}
				if !dp[i][j] {
					dp[i][j] = dp[i][j+2]
				}
			}
		}
	}
	return dp[0][0]
}

func initDpMap(s, e []byte) [][]bool {
	slen := len(s)
	elen := len(e)
	dp := make([][]bool, slen+1)
	for i := range dp {
		dp[i] = make([]bool, elen+1)
	}
	dp[slen][elen] = true
	for j := elen - 2; j > -1; j -= 2 {
		if e[j] != '*' && e[j+1] == '*' {
			dp[slen][j] = true
		} else {
			break
		}
	}
	if slen > 0 && elen > 0 {
		if e[elen-1] == '.' || s[slen-1] == e[slen-1] {
			dp[slen-1][elen-1] = true
		}
	}
	return dp
}

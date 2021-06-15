package string_type

func removeDuplicateLetters(s string) string {
	str := []byte(s)
	// 小写字母ASCII 码值范围为[97~122]，所以用长度为26 的数组做次数统计
	// 如果map[i] > -1，则代表ASCII 码值为i 的字符的出现次数
	// 如果map[i] == -1，则代表ASCII 码值为i 的字符不再考虑
	strMap := make([]int, 26)
	for i := 0; i < len(str); i++ {
		strMap[str[i]-'a']++
	}
	res := make([]byte, 26)
	index, L, R := 0, 0, 0
	for R != len(str) {
		// 如果当前字符是不再考虑的，直接跳过
		// 如果当前字符出现的次数减1 之后，后面还能出现，直接跳过
		tmp := strMap[str[R]-'a']
		if tmp == -1 || tmp > 0 {
			R++
		} else {
			// 当前字符需要考虑并且之后不会再出现
			// 在str[L...R]上所有需要考虑的字符中，找到ASCII码最小字符的位置
			pick := -1
			for i := L; i <= R; i++ {
				if strMap[str[i]-'a'] != -1 && (pick == -1 || str[i] < str[pick]) {
					pick = i
				}
			}
			// 把ASCII码最下的字符放到挑选的结果中
			res[index] = str[pick]
			index++
			// 在上一个的for 循环中，str[L..R]范围内每种字符出现的次数都减少了
			// 需要把str[pick + 1..R]中每种字符出现的次数加回来
			for i := pick + 1; i <= R; i++ {
				if strMap[str[i]-'a'] != -1 {
					strMap[str[i]-'a']++
				}
			}
			// 选出ASCII码最下的字符，以后不再考虑
			strMap[str[pick]-'a'] = -1
			// 继续在str[pick+1.......]上重复这个过程
			L = pick + 1
			R = L
		}
	}
	return string(res)
}

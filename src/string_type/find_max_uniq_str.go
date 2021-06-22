package string_type

import "algorithm/src/util"

/**
找到字符串的最长无重复字符子串
给定一个字符串str，返回str 的最长无重复字符子串的长度。
*/

func maxUnique(str string) int {
	if len(str) == 0 {
		return 0
	}
	chas := []byte(str)
	maps := make([]int, 256)
	for i := 0; i < 256; i++ {
		maps[i] = -1
	}
	lens, pre, cur := 0, -1, 0
	for i := 0; i != len(chas); i++ {
		pre = util.Max(pre, maps[chas[i]])
		cur = i - pre
		lens = util.Max(lens, cur)
		maps[chas[i]] = i
	}
	return lens
}

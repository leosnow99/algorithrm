package string_type

import (
	"algorithm/util"
	"math"
)

/**
最小包含子串的长度

给定字符串str1 和str2，求str1 的子串中含有str2 所有字符的最小子串长度。
*/

// 先通过right 向右扩，让所有 的字符被“有效”地还完，都还完时，
// 被框住的子串肯定是符合要求的，但还要经过left 向右缩的过程来看被框住的子串能不能变得更短。
func minLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 || len(str1) < len(str2) {
		return 0
	}
	chas1 := []byte(str1)
	chas2 := []byte(str2)
	maps := make([]int, 256)
	for _, c := range chas2 {
		maps[c]++
	}
	left, right, match, minLen := 0, 0, len(chas2), math.MaxInt32
	for right < len(chas1) {
		maps[chas1[right]]--
		if maps[chas1[right]] >= 0 {
			match--
		}
		if match == 0 {
			for maps[chas1[left]] < 0 {
				maps[chas1[left]]++
				left++
			}
			minLen = util.Min(minLen, right-left+1)
			match++
			maps[chas1[left]]++
			left++
		}
		right++
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

package string_type

/**
找到指定的新类型字符

新类型字符的定义如下：
1．新类型字符是长度为1 或者2 的字符串。
2．表现形式可以仅是小写字母，例如，"e"；也可以是大写字母+小写字母，例如，"Ab"；
还可以是大写字母+大写字母，例如，"DC"。
现在给定一个字符串str，str 一定是若干新类型字符正确组合的结果。比如"eaCCBi"，由新
类型字符"e"、"a"、"CC"和"Bi"拼成。再给定一个整数k，代表str 中的位置。请返回被k 位置指
定的新类型字符。
*/

func pointNewChar(s string, k int) string {
	if len(s) == 0 || k < 0 || k >= len(s) {
		return ""
	}
	chas := []byte(s)
	uNum := 0
	for i := k - 1; i >= 0; i-- {
		if chas[i] <= 'A' {
			break
		}
		uNum++
	}
	if uNum&1 == 1 {
		return string(chas[k-1 : k+1])
	}
	if chas[k] > 'A' {
		return string(chas[k : k+2])
	}
	return string(chas[k])
}

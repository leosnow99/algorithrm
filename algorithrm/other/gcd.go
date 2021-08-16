package other

/**
一行代码求两个数的最大公约数
【题目】
给定两个不等于0 的整数M 和N，求M 和N 的最大公约数。
*/

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	return gcd(n, m%n)
}

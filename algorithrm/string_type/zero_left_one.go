package string_type

// 给定一个整数N，求由"0"字符与"1"字符组成的长度为N 的所有字符串中，满足"0"字符的
// 左边必有"1"字符的字符串数量。

func getNum1(n int) int {
	if n < 1 {
		return 0
	}
	return processZeroLeftOne(1, n)
}

func processZeroLeftOne(i, n int) int {
	if i == n-1 {
		return 2
	}
	if i == n {
		return 1
	}
	return processZeroLeftOne(i+1, n) + processZeroLeftOne(i+2, n)
}

func getNum2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	pre, cur, tmp := 1, 1, 0
	for i := 2; i <= n; i++ {
		tmp = cur
		cur += pre
		pre = tmp
	}
	return cur

}

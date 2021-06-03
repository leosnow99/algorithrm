package dp

// 数字字符串转换为字母组合的种数

// 给定一个字符串str，str 全部由数字字符组成，如果str 中某一个或某相邻两个字符组成的
// 子串值在1~26 之间，则这个子串可以转换为一个字母。规定"1"转换为"A"，"2"转换为"B"，"3"
// 转换为"C"……"26"转换为"Z"。写一个函数，求str 有多少种不同的转换结果，并返回种数。

// 暴力递归
func num1(str string) int {
	if str == "" {
		return 0
	}
	chs := []byte(str)
	return processNum(chs, 0)
}

func processNum(chs []byte, i int) int {
	if i == len(chs) {
		return 1
	}
	if chs[i] == '0' {
		return 0
	}
	res := processNum(chs, i+1)
	if i+1 < len(chs) && (chs[i]-'0')*10+(chs[i+1]-'0') < 27 {
		res += processNum(chs, i+2)
	}
	return res
}

// 动态规范
func num2(str string) int {
	if str == "" {
		return 0
	}
	chs := []byte(str)
	cur := 1
	if chs[len(chs)-1] == '0' {
		cur = 0
	}
	next, tmp := 1, 0
	for i := len(chs) - 2; i >= 0; i-- {
		if chs[i] == '0' {
			next = cur
			cur = 0
		} else {
			tmp = cur
			if (chs[i]-'0')*10+chs[i+1]-'0' < 27 {
				cur += next
			}
			next = tmp
		}
	}
	return cur
}

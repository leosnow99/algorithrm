package dp

import (
	"algorithm/src/util"
)

/**
N 皇后问题

N 皇后问题是指在N×N 的棋盘上要摆N 个皇后，要求任何两个皇后不同行、不同列，也不
在同一条斜线上。给定一个整数n，返回n 皇后的摆法有多少种。
*/

func queenNum1(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return queenProcess1(0, record, n)
}

func queenProcess1(i int, record []int, n int) int {
	if i == n {
		return 0
	}
	res := 0
	for j := 0; j < n; j++ {
		if queenIsValid(record, i, j) {
			record[i] = j
			res += queenProcess1(i+1, record, n)
		}
	}
	return res
}

func queenIsValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || util.Abs(record[k]-j) == util.Abs(i-k) {
			return false
		}
	}
	return true
}

//使用了位运算来加速。具体加速的递归 过程中，找到每一行还有哪些位置可以放置皇后的判断过程。因为整个过程比较自然
func queenNum2(n int) int {
	// 因为本方法中位运算的载体是int 型变量，所以该方法只能算1~32 皇后问题
	// 如果想计算更多的皇后问题，需使用包含更多位的变量
	if n < 1 || n > 32 {
		return 0
	}
	var upperLim int
	if n == 32 {
		upperLim = -1
	} else {
		upperLim = 1<<n - 1
	}
	return queenProcess2(upperLim, 0, 0, 0)
}

func queenProcess2(upperLim, colLim, leftDiaLim, rightDiaLim int) int {
	if colLim == upperLim {
		return 1
	}
	pos := upperLim & (^(colLim | leftDiaLim | rightDiaLim))
	mostRightOne, res := 0, 0
	for ; pos != 0; {
		mostRightOne = pos & (^pos + 1)
		pos = pos - mostRightOne
		res += queenProcess2(upperLim, colLim|mostRightOne, (leftDiaLim|mostRightOne)<<1,
			(rightDiaLim|mostRightOne)>>1)
	}
	return res
}

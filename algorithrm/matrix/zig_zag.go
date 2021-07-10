package matrix

import "fmt"

/**
“之”字形打印矩阵
【题目】
给定一个矩阵matrix，按照“之”字形的方式打印这个矩阵
*/

func printMatrixZigZag(matrix [][]int) {
	tC, tR, dC, dR := 0, 0, 0, 0
	endR, endC := len(matrix)-1, len(matrix[0])-1
	fromUp := false
	for tR != endR+1 {
		printLevel(matrix, tR, tC, dR, dC, fromUp)
		if tC == endC {
			tR = tR + 1
		} else {
			tC++
		}
		if dR == endR {
			dC++
		} else {
			dR++
		}
		fromUp = !fromUp
	}
}

func printLevel(m [][]int, tR, tC, dR, dC int, f bool) {
	if f {
		for tR != dR+1 {
			fmt.Println(m[tR][tC], " ")
			tR++
			tC--
		}
	} else {
		for dR != tR-1 {
			fmt.Println(m[dR][dC], " ")
			dR--
			dC++
		}
	}
}

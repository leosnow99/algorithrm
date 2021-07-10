package matrix

import "fmt"

/**
转圈打印矩阵
【题目】
给定一个整型矩阵matrix，请按照转圈的方式打印它。
*/

func spiralOrderPrint(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	tR, tC := 0, 0
	dR, dC := len(matrix), len(matrix[0])
	for tR <= dR && tC <= dC {
		printEdge(matrix, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
}

func printEdge(matrix [][]int, tR, tC, dR, dC int) {
	if tR == dR {
		// 子矩阵只有一行
		for i := tC; i <= dC; i++ {
			fmt.Println(matrix[tR][i], " ")
		}
	} else if tC == dC {
		// 子矩阵只有一列
		for i := tR; i <= dR; i++ {
			fmt.Println(matrix[i][tC], " ")
		}
	} else {
		curC, curR := tC, tR
		for curC != dC {
			fmt.Println(matrix[tR][curC], " ")
			curC++
		}
		for curR != dR {
			fmt.Println(matrix[curR][dC], " ")
			curR++
		}
		for curC != tC {
			fmt.Println(matrix[dR][curC], " ")
			curC--
		}
		for curR != tR {
			fmt.Println(matrix[curR][tC], " ")
			curR--
		}
	}
}

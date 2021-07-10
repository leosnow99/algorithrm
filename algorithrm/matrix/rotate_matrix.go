package matrix

/**
将正方形矩阵顺时针转动90°
【题目】
给定一个N×N 的矩阵matrix，把这个矩阵调整成顺时针转动90°后的形式。
*/

func rotate(matrix [][]int) {
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

func rotateEdge(m [][]int, tR, tC, dR, dC int) {
	times := dC - tC
	for i := 0; i < times; i++ {
		m[tR][tC+i], m[tR+i][dC], m[dR][dC-i], m[dR-i][tC] = m[dR-i][tC], m[tR][tC+i], m[tR+i][dC], m[dR][tC-i]
	}
}

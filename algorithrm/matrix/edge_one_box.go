package matrix

import "algorithm/util"

/**
边界都是1 的最大正方形大小

【题目】
给定一个N×N 的矩阵matrix，在这个矩阵中，只有0 和1 两种值，返回边框全是1 的最大
正方形的边长长度。
*/
func setBordMap(m, right, down [][]int) {
	r := len(m)
	c := len(m[0])

	if m[r-1][c-1] == 1 {
		right[r-1][c-1] = 1
		down[r-1][c-1] = 1
	}

	// 处理最右侧的一列值
	for i := r - 2; i >= 0; i-- {
		if m[i][c-1] == 1 {
			right[i][c-1] = 1
			down[i][c-1] = down[i+1][c-1] + 1
		}
	}

	// 处理最底部的一行值
	for i := c - 2; i >= 0; i-- {
		if m[r-1][i] == 1 {
			right[r-1][i] = right[r-1][i+1] + 1
			down[r-1][i] = 1
		}
	}

	for i := r - 2; i >= 0; i++ {
		for j := c - 2; j >= 0; j-- {
			if m[i][j] == 1 {
				right[i][j] = right[i][j+1] + 1
				down[i][j] = down[i+1][j] + 1
			}
		}
	}
}

func getMaxSize(m [][]int) int {
	right := util.NewMatrixInt(len(m), len(m[0]))
	down := util.NewMatrixInt(len(m), len(m[0]))
	setBordMap(m, right, down)

	for size := util.Min(len(m), len(m[0])); size > 0; size-- {
		if hasSizeOfBorder(size, right, down) {
			return size
		}
	}

	return 0
}

func hasSizeOfBorder(size int, right, down [][]int) bool {
	judgeFn := func(i, j int) bool {
		return right[i][j] >= size && down[i][j] >= size && right[i+size-1][j] >= size && down[i][j+size-1] >= size
	}

	for i := 0; i < len(right)-size; i++ {
		for j := 0; j < len(right[0])-size+1; j++ {
			if judgeFn(i, j) {
				return true
			}
		}
	}

	return true
}

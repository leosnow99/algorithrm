package matrix

/**
在行列都排好序的矩阵中找指定数
【题目】
给定一个N×M 的整型矩阵matrix 和一个整数K，matrix 的每一行和每一列都是排好序的。
实现一个函数，判断K 是否在matrix 中。
*/

// 1．从矩阵最右上角的数开始寻找（row=0，col=M-1）。
//2．比较当前数matrix[row][col]与K 的关系：
// 如果与 K 相等，说明已找到，直接返回true。
// 如果比 K 大，因为矩阵每一列都已排好序，所以在当前数所在的列中，处于当前数下
//方的数都会比K 大，则没有必要继续在第col 列上寻找，令col=col-1，重复步骤2。
// 如果比 K 小，因为矩阵每一行都已排好序，所以在当前数所在的行中，处于当前数左
//方的数都会比K 小，则没有必要继续在第row 行上寻找，令row=row+1，重复步骤2。
//3．如果找到越界都没有发现与K 相等的数，则返回false。
func isContains(matrix [][]int, k int) bool {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix)-1 && col > -1 {
		if matrix[row][col] == k {
			return true
		} else if matrix[row][col] < k {
			col++
		} else {
			row++
		}
	}

	return false
}

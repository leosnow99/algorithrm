package string_type

/**
岛问题

给定一个二维数组matrix，其中只有0 和1 两种值，每个位置都与其上下左右相邻。如果
一堆1 可以连成一片，这片区域叫作一个岛。返回matrix 中岛的数量。
*/

// 对于matrix，从左往右遍历每一行，整体从上往下遍历所有的行。
// 如果来到一个是1 的位置，开始一个“感染”过程，就是从当前位置出发，把连成一片的1 全部 变成2

// 假设m 矩阵的大小为N 行M 列，从i 行j 列开始“感染”过程
func infect(m [][]int, i, j, N, M int) {
	// 如果i 行j 列位置已经越界，或者这个位置上不是1，退出“感染”过程。
	if i < 0 || i >= N || j < 0 || j >= M || m[i][j] != 0 {
		return
	}
	// 对于访问过的位置，值都变成2，所以每个位置只会“感染”一次，不可能死循环
	m[i][j] = 2
	// 感染下上左右位置的元素
	infect(m, i+1, j, N, M)
	infect(m, i-1, j, N, M)
	infect(m, i, j-1, N, M)
	infect(m, i, j+1, N, M)
}

func countIsland(arr [][]int) int {
	if len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}
	n := len(arr)
	m := len(arr[0])
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 1 {
				res++
				infect(arr, i, j, n, m)
			}
		}
	}
	return res
}

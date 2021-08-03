package matrix

import (
	"algorithm/util"

	"github.com/ahrtr/gocontainer/list"
)

/**
求最短通路值
【题目】
用一个整型矩阵matrix 表示一个网络，1 代表有路，0 代表无路，每一个位置只要不越界，
都有上下左右4 个方向，求从最左上角到最右下角的最短通路值。
*/

/**
使用宽度优先遍历即可，如果矩阵大小为N×M，时间复杂度为O(N×M)

1．开始时生成map 矩阵，map[i][j]的含义是从（0,0）位置走到（i,j）位置最短的路径值。
	然后将左上角位置（0,0）的行坐标与列坐标放入行队列rQ 和列队列cQ。
2．不断从队列弹出一个位置（r,c），然后看这个位置的上下左右四个位置哪些在matrix 上的值是1，这些都是能走的位置。
3．将那些能走的位置设置好各自在map 中的值，即map[r][c]+1。同时将这些位置加入rQ和cQ 中，用队列完成宽度优先遍历。
4．在步骤3 中，如果一个位置之前走过，就不要重复走，这个逻辑可以根据一个位置在map 中的值来确定，比如map[i][j]!=0，
	就可以知道这个位置之前已经走过。
5．一直重复步骤2～步骤4。直到遇到右下角位置，说明已经找到终点，返回终点在map中的值即可，
	如果rQ 和cQ 已经为空都没有遇到终点位置，说明不存在这样一条路径，返回0。

每个位置最多走一遍，所以时间复杂度为O(N×M)、额外空间复杂度也是O(N×M)。
*/

func minPathValue(m [][]int) int {
	if len(m) == 0 || len(m[0]) == 0 || m[0][0] == 0 || m[len(m)-1][len(m[0])-1] == 0 {
		return 0
	}

	maps := util.NewMatrixInt(len(m), len(m[0]))
	maps[0][0] = 1

	rQ := list.NewLinkedList()
	cQ := list.NewLinkedList()
	rQ.Add(0)
	cQ.Add(0)

	r, c := 0, 0
	for !rQ.IsEmpty() {
		get, _ := rQ.Get(0)
		r = get.(int)

		get, _ = cQ.Get(0)
		c = get.(int)

		if r == len(m)-1 && c == len(m[0])-1 {
			return maps[r][c]
		}

		walkTo(maps[r][c], r-1, c, m, maps, rQ, cQ) // 上
		walkTo(maps[r][c], r+1, c, m, maps, rQ, cQ) // 下
		walkTo(maps[r][c], r, c-1, m, maps, rQ, cQ) // 左
		walkTo(maps[r][c], r, c+1, m, maps, rQ, cQ) // 右
	}

	return 0
}

func walkTo(pre, toR, toC int, m, maps [][]int, rQ, cQ list.Interface) {
	if toR < 0 || toR == len(m) || toC < 0 || toC == len(m[0]) || m[toR][toC] != 1 || maps[toR][toC] != 0 {
		return
	}

	maps[toR][toC] = pre + 1
	rQ.Add(toR)
	cQ.Add(toC)
}

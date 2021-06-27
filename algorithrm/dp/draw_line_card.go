package dp

import "algorithm/util"

/**
排成一条线的纸牌博弈问题

给定一个整型数组arr，代表数值不同的纸牌排成一条线。玩家A 和玩家B 依次拿走每张纸
牌，规定玩家A 先拿，玩家B 后拿，但是每个玩家每次只能拿走最左或最右的纸牌，玩家A 和
玩家B 都绝顶聪明。请返回最后获胜者的分数。
*/
func win1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return util.Max(f(arr, 0, len(arr)-1), s(arr, 0, len(arr)-1))
}

func f(arr []int, i, j int) int {
	if i == j {
		return arr[i]
	}
	return util.Max(arr[i]+s(arr, i+1, j), s(arr, i, j-1))
}

func s(arr []int, i, j int) int {
	if i == j {
		return 0
	}
	return util.Min(f(arr, i+1, j), f(arr, i, j-1))
}

func win2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	f := make([][]int, len(arr))
	s := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		f[i] = make([]int, len(arr))
		s[i] = make([]int, len(arr))
	}
	for i := 0; i < len(arr); i++ {
		f[i][i] = arr[i]
		for j := i - 1; j >= 0; j-- {
			f[j][i] = util.Max(arr[j]+s[j+1][i], arr[j]+s[j][i-1])
			s[i][j] = util.Min(f[j+1][i], f[j][i-1])
		}
	}
	return util.Max(f[0][len(arr)-1], s[0][len(arr)-1])
}

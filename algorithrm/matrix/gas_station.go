package matrix

/**
加油站良好出发点问题
【题目】
N 个加油站组成一个环形，给定两个长度都是N 的非负数组oil 和dis（N>1），oil[i]代表第
i 个加油站存的油可以跑多少千米，dis[i]代表第i 个加油站到环中下一个加油站相隔多少千米。
假设你有一辆油箱足够大的车，初始时车里没有油。如果车从第i 个加油站出发，最终可以回
到这个加油站，那么第i 个加油站就算良好出发点，否则就不算。请返回长度为N 的boolean
型数组res，res[i]代表第i 个加油站是不是良好出发点。
*/

func station(dis, oil []int) []bool {
	if len(dis) < 2 || len(dis) != len(oil) {
		return nil
	}

	// 良好出发点存在的前提：init 点存在。如果 init 点不存在，说明所有的加油站纯能值都小于 0，必然所有点都不是良好出发点。
	init := changeDisArrayGetInit(dis, oil)
	if init == -1 {
		return make([]bool, len(dis), len(dis))
	}

	return enlargeArea(dis, init)
}

// 转换 dis 数组为「纯能数组」，并且返回一个大于等于 0 的位置。
func changeDisArrayGetInit(dis, oil []int) int {
	if len(dis) == 0 || len(dis) != len(oil) {
		return 0
	}

	var initPoint = -1
	for idx, o := range oil {
		dis[idx] = o - dis[idx]

		if initPoint < 0 && initPoint < dis[idx] {
			initPoint = dis[idx]
		}
	}

	return initPoint
}

// 扩充连通区
func enlargeArea(dis []int, init int) []bool {
	res := make([]bool, len(dis), len(dis))
	start := init
	end := nextIndex(start, len(dis))
	// need 值为从 start 位置顺时针扩充连通区的要求
	need := 0
	// rest 值为从 end 位置逆时针扩充连通区的资源
	rest := 0

	for {
		// 当前来到的start 已经在连通区域中，可以确定后续的开始点一定无法转完一圈
		if start != init && start == lastIndex(end, len(dis)) {
			break
		}

		// 当前来到的start 不在连通区域中，就扩充连通区域
		if dis[start] < need { // 无法到达
			need -= dis[start]
		} else {
			rest += dis[start] - need
			need = 0
			for rest >= 0 && end != start {
				rest += dis[end]
				end = nextIndex(end, len(dis))
			}
			// 如果连通区域已经覆盖整个环，当前的start 是良好出发点，进入2 阶段
			if rest >= 0 {
				res[start] = true
				connectGood(dis, lastIndex(start, len(dis)), init, res)
				break
			}
		}

		start = lastIndex(end, len(dis))
		if start == init {
			break
		}
	}

	return res
}

// 已知start 的next 方向上有一个良好出发点
// start 如果可以达到这个良好出发点，那么从start 出发一定可以转一圈
func connectGood(dis []int, start, init int, res []bool) {
	need := 0
	for start != init {
		if dis[start] < need {
			need -= dis[start]
		} else {
			res[start] = true
			need = 0
		}
		start = lastIndex(start, len(dis))
	}
}

func nextIndex(index, size int) int {
	if index == size-1 {
		return 0
	}

	return index + 1
}

func lastIndex(index, size int) int {
	if index == 0 {
		return size
	}
	return index - 1
}

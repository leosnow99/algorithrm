package dp

import "algorithm/src/util"

/**
跳跃游戏

给定数组arr，arr[i]==k 代表可以从位置i 向右跳1~k 个距离。比如，arr[2]==3，代表可以
从位置2 跳到位置3、位置4 或位置5。如果从位置0 出发，返回最少跳几次能跳到arr 最后的
位置上
 */

func jumpGame(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	jump, cur, next := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		if cur < i {
			jump++
			cur = next
		}
		next = util.Max(next, i+arr[i])
	}
	return jump
}

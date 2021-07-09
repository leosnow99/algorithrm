package bit

/**
在其他数都出现k 次的数组中找到只出现一次的数
【题目】
给定一个整型数组arr 和一个大于1 的整数k。已知arr 中只有1 个数出现了1 次，其他的
数都出现了k 次，请返回只出现了1 次的数。
【要求】
时间复杂度为O(N)，额外空间复杂度为O(1)。
*/

// 首先设置一个变量eO，它是一个32 位的k 进制数，且每个位置上都是0。然后遍历arr，把遍历到的每一个整数都转换为k 进制数，
// 然后与eO 进行无进位相加。遍历结束时，把32 位的k 进制数eORes 转换为十进制整数，就是我们想要的结果。
// 因为k 个相同的k 进制数无进位相加，结果一定是每一位上都是0 的k 进制数，所以只出现一次的那个数最终就会剩下来。
func onceNum(arr []int, k int) int {
	eO := make([]int, 32)
	for _, num := range arr {
		setExclusiveOr(eO, num, k)
	}
	return getNumFromKSysNum(eO, k)
}

func setExclusiveOr(eO []int, value, k int) {
	curKSysNum := getKSysNumFromNum(value, k)
	for i := 0; i != len(eO); i++ {
		eO[i] = (eO[i] + curKSysNum[i]) % k
	}
}

func getKSysNumFromNum(value, k int) []int {
	res := make([]int, 32)
	index := 0
	for value != 0 {
		res[index] = value % k
		index++
		value = value / k
	}
	return res
}

func getNumFromKSysNum(eO []int, k int) int {
	res := 0
	for i := len(eO) - 1; i >= 0; i-- {
		res = res*k + eO[i]
	}

	return res
}

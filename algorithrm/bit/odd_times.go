package bit

/**
在其他数都出现偶数次的数组中找到出现奇数次的数
【题目】
给定一个整型数组arr，其中只有一个数出现了奇数次，其他的数都出现了偶数次，打印这
个数。
进阶问题：有两个数出现了奇数次，其他的数都出现了偶数次，打印这两个数。
【要求】
时间复杂度为O(N)，额外空间复杂度为O(1)。
*/

func printOddNum(arr []int) int {
	eO := 0
	for _, num := range arr {
		eO ^= num
	}
	return eO
}

// 如果只有A 和B 出现了奇数次，那么最后的异或结果eO 就是A^B。所以，如果数组中有
// 两个出现了奇数次的数，最终的eO 一定不等于0。那么肯定能在32 位整数eO 上找到一个不等
// 于0 的bit 位，假设是第k 位不等于0。eO 在第k 位不等于0，说明A 和B 的第k 位肯定一个是
// 1，另一个是0。接下来再设置一个变量记为eOhasOne，然后遍历一次数组。在这次遍历时，
// eOhasOne 只与第k 位上是1 的整数异或，其他的数忽略。那么在第二次遍历结束后，eOhasOne
// 就是A 或者B 中的一个， 而eO^eOhasOne 就是另外一个出现奇数次的数
func printOddNum2(arr []int) (int, int) {
	eO, eOhasOne := 0, 0
	for _, num := range arr {
		eO ^= num
	}
	rightOne := eO & (eO + 1)
	for _, num := range arr {
		if num&rightOne != 0 {
			eOhasOne ^= num
		}
	}
	return eOhasOne, eO ^ eOhasOne
}

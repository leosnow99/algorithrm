package bit

/**
不用额外变量交换两个整数的值
如何不用任何额外变量交换两个整数的值？
*/

func swapTwoInt(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
}

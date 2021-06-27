package bit

/**
不用做任何比较判断找出两个数中较大的数
【题目】
给定两个32 位整数a 和b，返回a 和b 中较大的。
【要求】
不用做任何比较判断。
*/

// flip 函数的功能是如果 n 为1，返回0，如果 n 为0，返回1
func flip(n int) int {
	return n ^ 1
}

// sign 函数的功能是返回整数n 的符号，正数和 0 返回1，
func sign(n int) int {
	return flip((n >> 31) & 1)
}

// 方法一是有局限性的，那就是如果a-b 的值出现溢出，返回结果就不正确。
func getMax1(a, b int) int {
	c := a - b
	scA := sign(c)
	scB := flip(sign(c))
	return scA*a + scB*b
}

// 如果a 的符号与b 的符号不同（difSab==1，sameSab==0），则有：
// 	 如果 a 为0 或正，那么b 为负（sa==1，sb==0），应该返回a；
// 	 如果 a 为负，那么b 为0 或正（sa==0，sb==1），应该返回b。
// 如果a 的符号与b 的符号相同（difSab==0，sameSab==1），这种情况下，a-b 的值绝对不会溢出：
// 	 如果 a-b 为0 或正（sc==1），返回a；
// 	 如果 a-b 为负（sc==0），返回b；
// 综上所述，应该返回a * (difSab * sa + sameSab * sc) + b * flip(difSab * sa + sameSab * sc)。
func getMax2(a, b int) int {
	c := a - b
	scA := sign(a)
	scB := sign(b)
	scC := sign(c)
	difSab := scA ^ scB
	sameSab := flip(difSab)
	returnA := difSab * scA + sameSab * scC
	returnB := flip(returnA)
	return returnA * a + returnB * b
}

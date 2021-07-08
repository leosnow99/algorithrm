package bit

import "math"

func add(a, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}

	return sum
}

func negNum(n int) int {
	return add(^n, 1)
}

func minus(a, b int) int {
	return add(a, negNum(b))
}

// 用位运算实现乘法运算。a×b 的结果可以写成a×20×b0+a×21×b1+…+a×2i×bi+…+ a×231×b31，
// 其中，bi 为0 或1 代表整数b 的二进制数表达中第i 位的值
func multi(a, b int) int {
	res := 0
	for b != 0 {
		if (b & 1) != 0 {
			res = add(res, a)
		}
		a <<= 1
		b >>= 1
	}
	return res
}

func isNeg(n int) bool {
	return n < 0
}

func div(a, b int) int {
	x := a
	if isNeg(a) {
		x = negNum(a)
	}

	y := b
	if isNeg(b) {
		y = negNum(b)
	}

	res := 0
	for i := 31; i >= 0; i-- {
		if (x >> i) > y {
			res |= 1 << i
			x = minus(x, y<<i)
		}
	}
	if !(isNeg(a) && isNeg(b)) && isNeg(a) || isNeg(b) {
		return negNum(res)
	}
	return res
}

// 除法实现还剩非常关键的最后一步。以上方法可以算绝大多数的情况，但我们知道32 位整
// 数的最小值为-2 147 483 648，最大值为2 147 483 647，最小值的绝对值比最大值的绝对值大1，
// 所以，如果a 或b 等于最小值，是转不成相对应的正数的。可以总结一下：
//	 如果 a 和b 都不为最小值，直接使用以上过程，返回div(a,b)。
//	 如果 a 和b 都为最小值，a/b 的结果为1，直接返回1。
//	 如果 a 不为最小值，而b 为最小值，a/b 的结果为0，直接返回0。
//	 如果 a 为最小值，而b 不为最小值，怎么办？
// 第1～3 种情况处理都比较容易，对于情况4 就棘手很多。我们举个简单的例子说明本书是
// 如何处理这种情况的。为了方便说明，我们假设整数的最大值为9，而最小值为-10。当a 和b
// 属于[0,9]的范围时，我们可以正确地计算a/b。当a 和b 都属于[-9,9]时，我们可以计算，也就
// 是情况1；当a 和b 都等于-10 时，我们也可以计算，就是情况2；当a 属于[-9,9]，而b 等于-10
// 时，我们也能计算，就是情况3；当a 等于-10，而b 属于[-9,9]时，如何计算呢？
//	1．假设a=-10，b=5。
//	2．计算(a+1)/b 的结果，记为c。对本例来讲就是-9/5 的结果，c=-1。
//	3．计算c×b 的结果。对本例来讲，-1×5=-5。
// 	4．计算a-(c×b)，即-10-(-5)=-5。
// 	5．计算(a-(c×b))/b 的结果，记为rest，意义是修正值，即-5/5=-1。
//	6．返回c+rest 的结果。
// 也就是说，既然我们对最小值无能为力，那么就把最小值增加一点，计算出一个结果，然
// 后根据这个结果再修正一下，得到最终的结果。
func divide(a, b int) int {
	if b == 0 {
		panic("divisor is 0")
	}
	if a == math.MinInt32 && b == math.MinInt32 {
		return 1
	} else if b == math.MinInt32 {
		return 0
	} else if a == math.MinInt32 {
		res := div(add(a, 1), b)
		return add(res, div(minus(a, multi(res, b)), b))
	}

	return div(a, b)
}

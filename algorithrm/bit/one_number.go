package bit

/**
整数的二进制数表达中有多少个1
【题目】
给定一个32 位整数n，可为0，可为正，也可为负，返回该整数二进制数表达中1 的个数。
*/

func count1(n int) int {
	var res = 0
	for n != 0 {
		res += n & 1
		n = int(uint32(n) >> 1)
	}
	return res
}

// 每次进行n&=(n-1)操作时，在 for 循环中就可以忽略掉bit 位上为0 的部分。
func count2(n int) int {
	var res int
	for n != 0 {
		n &= n - 1
		res++
	}
	return res
}

// n & (~n + 1)的含义是得到n 中最右侧的1
func count3(n int) int {
	var res int
	for n != 0 {
		n -= n & (^n + 1)
		res++
	}
	return res
}

func count4(n int) int {
	n = (n & 0x55555555) + (int(uint32(n)>>1) & 0x55555555)
	n = (n & 0x33333333) + (int(uint32(n)>>1) & 0x33333333)
	n = (n & 0x0f0f0f0f) + (int(uint32(n)>>1) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + (int(uint32(n)>>1) & 0x00ff00ff)
	n = (n & 0x0000ffff) + (int(uint32(n)>>1) & 0x0000ffff)
	return n
}

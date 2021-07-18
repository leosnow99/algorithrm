package matrix

/**
奇数下标都是奇数或者偶数下标都是偶数
【题目】
	给定一个长度不小于2 的数组arr，实现一个函数调整arr，要么让所有的偶数下标都是偶数，要么让所有的奇数下标都是奇数。
【要求】
	如果 arr 的长度为N，函数要求时间复杂度为O(N)、额外空间复杂度为O(1)。
*/

// 1．设置变量even，表示目前arr 最左边的偶数下标，初始时even=0。
// 2．设置变量odd，表示目前arr 最左边的奇数下标，初始时odd=1。
// 3．不断检查arr 的最后一个数，即arr[N-1]。如果arr[N-1]是偶数，交换arr[N-1]和arr[even]，
//   然后令even=even+2。如果arr[N-1]是奇数，交换arr[N-1]和arr[odd]，然后令odd=odd+2。继续重复步骤3。
// 4．如果even 或者odd 大于或等于N，过程停止。
func modifyOddEven(arr []int) {
	if len(arr) < 2 {
		return
	}

	even, odd, end := 0, 1, len(arr)-1

	for even < end || odd < end {
		if arr[end]&1 == 0 {
			arr[end], arr[even] = arr[even], arr[end]
			even += 2
		} else {
			arr[end], arr[odd] = arr[odd], arr[even]
			odd += 2
		}
	}
}

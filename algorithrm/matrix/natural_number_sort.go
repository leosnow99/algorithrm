package matrix

/**
自然数数组的排序
【题目】
	给定一个长度为N 的整型数组arr，其中有N 个互不相等的自然数1～N。请实现arr 的排序，但是不要把下标0～N-1 位置上的数通过直接赋值的方式替换成1～N。
【要求】
	时间复杂度为O(N)，额外空间复杂度为O(1)。

 arr 在调整之后应该是下标从0 到N-1 的位置上依次放着1~N，即arr[index]=index+1。
*/

// 1．从左到右遍历arr，假设当前遍历到i 位置。
// 2．如果arr[i]==i+1，说明当前的位置不需要调整，继续遍历下一个位置。
// 3．如果arr[i]!=i+1，说明此时i 位置的数arr[i]不应该放在i 位置上，接下来将进行跳的过程。
func natureSort1(arr []int) {
	tmp, next := 0, 0

	for i := 0; i != len(arr); i++ {
		tmp = arr[i]
		for arr[i] != i+1 {
			next = arr[tmp-1]
			arr[tmp-1] = tmp
			tmp = next
		}
	}
}

// 1．从左到右遍历arr，假设当前遍历到i 位置。
// 2．如果arr[i]==i+1，说明当前的位置不需要调整，继续遍历下一个位置。
// 3．如果arr[i]!=i+1，说明此时i 位置的数arr[i]不应该放在i 位置上，接下来将在i 位置进行交换过程。
func natureSort2(arr []int) {
	for i := 0; i != len(arr); i++ {
		for arr[i] != i+1 {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
	}
}

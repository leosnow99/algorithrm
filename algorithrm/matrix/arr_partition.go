package matrix

/**
数组的 partition 调整
【题目】
	给定一个有序数组arr，调整arr 使得这个数组的左半部分没有重复元素且升序，而不用保
	证右部分是否有序。
	例如，arr=[1,2,2,2,3,3,4,5,6,6,7,7,8,8,8,9]，调整之后arr=[1,2,3,4,5,6,7,8,9,…]。
	补充问题：给定一个数组arr，其中只可能含有0、1、2 三个值，请实现arr 的排序。
	另一种问法为：有一个数组，其中只有红球、蓝球和黄球，请实现红球全放在数组的左边，蓝球放在中间，黄球放在右边。
	另一种问法为：有一个数组，再给定一个值k，请实现比k 小的数都放在数组的左边，等
	于k 的数都放在数组的中间，比k 大的数都放在数组的右边。
【要求】
	1．所有题目实现的时间复杂度为O(N)。
	2．所有题目实现的额外空间复杂度为O(1)
*/

/**
1．生成变量u，含义是在arr[0..u]上都是无重复元素且升序的。也就是说，u 是这个区域最后的位置，初始时u=0，这个区域记为A。
2．生成变量i，利用i 做从左到右的遍历，在arr[u+1..i]上是不保证没有重复元素且升序的区域，i 是这个区域最后的位置，初始时i=1，这个区域记为B。
3．i 向右移动（i++）。因为数组整体有序，所以，如果arr[i]!=arr[u]，说明当前数arr[i]应该加入到A 区域里，
	交换arr[u+1]和arr[i]，此时A 的区域增加一个数（u++）；如果arr[i]==arr[u]，说明当前数arr[i]的值之前已经加入A 区域，此时不用再加入。
4．重复步骤3，直到所有的数遍历完。
*/
func leftUniq(arr []int) {
	if len(arr) < 2 {
		return
	}

	u, i := 0, 1
	for i != len(arr) {
		if arr[i] != arr[u] {
			u++
			arr[u], arr[i-1] = arr[i-1], arr[u]
		}
		i++
	}
}

/**
1．生成变量left，含义是在arr[0..left]（左区）上都是0，left 是这个区域当前最右的位置，初始时left 为-1。
2．生成变量index，利用这个变量做从左到右的遍历，含义是在arr[left+1..index]（中区）上都是1，
	index 是这个区域当前最右的位置，初始时index 为0。
3．生成变量right，含义是在arr[right..N-1]（右区）上都是2，right 是这个区域当前最左的位置，初始时right 为N。
4．index 表示遍历到arr 的一个位置：
	1）如果arr[index]==1，这个值应该直接加入到中区，index++之后重复步骤4。
	2）如果arr[index]==0，这个值应该加入到左区，arr[left+1]是中区最左的位置，所以把arr[index]和arr[left+1]交换之后，
		左区就扩大了，index++之后重复步骤4。
	3）如果arr[index]==2，这个值应该加入到右区，arr[right-1]是右区最左边数的左边，但也不属于中区，总之，在中区和右区的中间部分。
		把arr[index]和arr[right-1]交换之后，右区就向左扩大了（right--），但是此时arr[index]上的值未知，所以index 不变，重复步骤4。
5．当index==right 时，说明中区和右区成功对接，三个区域都划分好后，过程停止。在遍历中的每一步，要么index 增加，
	要么right 减少。如果index==right，过程就停止，所以时间复杂度就是O(N).
*/

func partitionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, index, right := -1, 0, len(arr)
	for index < right {
		if arr[index] == 0 {
			// arr[left++] 只有可能是 1 或 0
			left++
			arr[left], arr[index] = arr[index], arr[left]
			index++
		} else if arr[index] == 2 {
			right--
			arr[right], arr[index] = arr[index], arr[right]
		} else {
			index++
		}
	}
}

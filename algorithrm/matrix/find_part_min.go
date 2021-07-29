package matrix

/**
在数组中找到一个局部最小的位置

【题目】
定义局部最小的概念:
	arr 长度为1 时，arr[0]是局部最小。arr 的长度为N（N>1）时，如果arr[0]<arr[1]，那么arr[0]是局部最小；
	如果arr[N-1]<arr[N-2]，那么arr[N-1]是局部最小；
	如果0<i<N-1，既有arr[i]<arr[i-1]，又有arr[i]<arr[i+1]，那么arr[i]是局部最小。

给定无序数组arr，已知arr 中任意两个相邻的数都不相等。写一个函数，只需返回arr 中任意一个局部最小出现的位置即可。
*/

/**
本题可以利用二分查找做到时间复杂度为O(logN)、额外空间复杂度为O(1)，步骤如下：
	1．如果arr 为空或者长度为0，返回-1 表示不存在局部最小。
	2．如果arr 长度为1 或者arr[0]<arr[1]，说明arr[0]是局部最小，返回0。
	3．如果arr[N-1]<arr[N-2]，说明arr[N-1]是局部最小，返回N-1。
	4．如果arr 长度大于2 且arr 的左右两头都不是局部最小，则令left=1，right=N-2，然后进入步骤5 做二分查找。
	5．令mid=(left+right)/2，然后进行如下判断：
		1）如果arr[mid]>arr[mid-1]，可知在arr[left..mid-1]上肯定存在局部最小，令right=mid-1，重复步骤5。
		2）如果不满足1)，但arr[mid]>arr[mid+1]，可知在arr[mid+1..right]上肯定存在局部最小，令left=mid+1，重复步骤5。
		3）如果既不满足1)，也不满足2)，那么arr[mid]就是局部最小，直接返回mid。
	6．步骤5 一直进行二分查找，直到left==right 时停止，返回left 即可。
*/
func findMinIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}

	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}

	left, right, mid := 1, len(arr)-2, 0
	for left < right {
		mid = (left + right) / 2

		if arr[mid] > arr[mid-1] {
			right = mid - 1
			continue
		}
		if arr[mid] > arr[mid+1] {
			left = mid + 1
			continue
		}

		return mid
	}

	return left
}

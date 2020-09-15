package sort

//选择排序
//选择排序(Selection-sort)是一种简单直观的排序算法。它的工作原理：
//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
//然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
//以此类推，直到所有元素均排序完毕。
func SelectSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

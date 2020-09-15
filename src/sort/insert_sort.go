package sort

//直接插入排序
//插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。它
//的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
func InsertSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}

	}
}

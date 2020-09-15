package sort

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mergeProcess(arr, 0, len(arr)-1)
}

func mergeProcess(arr []int, left int, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)>>1
	mergeProcess(arr, left, mid)
	mergeProcess(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	i := 0
	leftIndex := left
	rightIndex := mid + 1
	for leftIndex <= mid && rightIndex <= right {
		if arr[leftIndex] < arr[rightIndex] {
			temp[i] = arr[leftIndex]
			leftIndex++
		} else {
			temp[i] = arr[rightIndex]
			rightIndex++
		}
		i++
	}

	for leftIndex <= mid {
		temp[i] = arr[leftIndex]
		i++
		leftIndex++
	}

	for rightIndex <= right {
		temp[i] = arr[rightIndex]
		i++
		rightIndex++
	}
	for i := 0; i < len(temp); i++ {
		arr[i+left] = temp[i]
	}
}

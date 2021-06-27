package sort

func HeapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	size := len(arr)
	size--
	arr[0], arr[size] = arr[size], arr[0]
	for size > 0 {
		heapify(arr, 0, size)
		size--
		arr[0], arr[size] = arr[size], arr[0]
	}

}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

func heapify(arr []int, index, size int) {
	left := index*2 + 1
	for left < size {
		largest := left
		if (left+1) < size && arr[left+1] > arr[left] {
			largest = left + 1
		}
		if largest == index {
			break
		}
		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
		left = index*2 + 1
	}

}

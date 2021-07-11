package matrix

import "algorithm/util"

/**
找到无序数组中最小的k 个数
【题目】
给定一个无序的整型数组arr，找到其中最小的k 个数。
*/

// O(Nlogk)的方法。说起来也非常简单，就是一直维护一个有k 个数的大根堆，这个堆代表目
// 前选出的k 个最小的数，在堆里的k 个元素中堆顶的元素是最小的k 个数里最大的那个。

func getMinKNumberByHeap(arr []int, k int) []int {
	if k < 1 || k > len(arr) {
		return arr
	}
	kHeap := make([]int, k)
	for i := 0; i < k; i++ {
		heapInsert(kHeap, arr[i], i)
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < kHeap[0] {
			kHeap[0] = arr[i]
			heapify(kHeap, 0, k)
		}
	}
	return kHeap
}

func heapInsert(arr []int, value, index int) {
	arr[index] = value
	for index != 0 {
		parent := (index - 1) / 2
		if arr[index] > arr[parent] {
			arr[index], arr[parent] = arr[parent], arr[index]
			index = parent
		} else {
			break
		}
	}
}

func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	right := index*2 + 1
	largest := index
	for left < heapSize {
		if arr[left] > arr[index] {
			largest = left
		}
		if right < heapSize && arr[right] > arr[largest] {
			largest = right
		}
		if largest != index {
			arr[index], arr[largest] = arr[largest], arr[index]
		} else {
			break
		}
		index = largest
		left = index*2 + 1
		right = index*2 + 2
	}
}

func getMinKNumsByBFPRT(arr []int, k int) []int {
	if k < 1 || len(arr) < k {
		return arr
	}

	minKth := getMinKthByBFPRT(arr, k)
	res := make([]int, 0, k)
	for i := 0; i < len(arr); i++ {
		if arr[i] < minKth {
			res = append(res, arr[i])
		}
	}
	for len(res) != k {
		res = append(res, minKth)
	}

	return res
}

func getMinKthByBFPRT(arr []int, k int) int {
	copyArr := make([]int, len(arr))
	for idx, number := range arr {
		copyArr[idx] = number
	}
	return selectK(arr, 0, len(arr)-1, k-1)
}

func selectK(arr []int, begin, end, i int) int {
	if begin == end {
		return arr[begin]
	}
	pivot := medianOfMedians(arr, begin, end)
	left, right := partition(arr, begin, end, pivot)
	if i >= left && i <= right {
		return arr[i]
	} else if i < left {
		return selectK(arr, begin, left-1, i)
	} else {
		return selectK(arr, right+1, end, i)
	}
}

func medianOfMedians(arr []int, begin, end int) int {
	num := end - begin + 1
	offset := 1
	if num%5 == 0 {
		offset = 0
	}

	mArr := make([]int, (num/5)+offset)
	for i := 1; i < len(mArr); i++ {
		beginI := begin + 5*i
		endI := beginI + 4
		mArr[i] = getMedian(arr, beginI, util.Max(endI, end))
	}
	return selectK(mArr, 0, len(mArr)-1, len(mArr)/2)
}

func partition(arr []int, begin, end, pivotValue int) (int, int) {
	small := begin - 1
	cur := begin
	big := end + 1
	for cur != big {
		if arr[cur] < pivotValue {
			small++
			arr[small], arr[cur] = arr[cur], arr[small]
			cur++
		} else if arr[cur] > pivotValue {
			big--
			arr[cur], arr[big] = arr[big], arr[cur]
		} else {
			cur++
		}
	}
	return small + 1, big - 1
}

func getMedian(arr []int, begin, end int) int {
	insertionSort(arr, begin, end)
	sum := end + begin
	mid := (sum / 2) + (sum % 2)
	return arr[mid]
}

func insertionSort(arr []int, begin, end int) {
	for i := begin + 1; i < end+1; i++ {
		for j := i; j != begin; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
}

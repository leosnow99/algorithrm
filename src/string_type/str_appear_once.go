package string_type

// 判断字符数组中是否所有的字符都只出现过一次

// 实现时间复杂度为O(N)的方法。
func isUnique1(chas []byte) bool {
	if len(chas) <= 1 {
		return true
	}
	maps := make([]bool, 256)
	for i := 0; i < 256; i++ {
		if maps[chas[i]] {
			return false
		}
		maps[chas[i]] = true
	}
	return true
}

// 在保证额外空间复杂度为O(1)的前提下，请实现时间复杂度尽量低的方法。
func isUnique2(chas []byte) bool {
	if len(chas) == 0 {
		return true
	}
	headSort(chas)
	for i := 1; i < len(chas); i++ {
		if chas[i] == chas[i-1] {
			return false
		}
	}
	return true
}

func headSort(chas []byte) {
	for i := 0; i < len(chas); i++ {
		heapInsert(chas, i)
	}
	for i := len(chas) - 1; i > 0; i-- {
		chas[0], chas[i] = chas[i], chas[0]
		heapify(chas, 0, i)
	}
}

func heapInsert(chas []byte, i int) {
	parent := 0
	for ; i != 0; {
		parent = (i - 1) >> 1
		if chas[parent] < chas[i] {
			chas[i], chas[parent] = chas[parent], chas[i]
			i = parent
		} else {
			break
		}
	}
}

func heapify(chas []byte, i, size int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	for left < size {
		if chas[left] > chas[i] {
			largest = left
		}
		if right < size && chas[right] > chas[largest] {
			largest = right
		}
		if largest != i {
			chas[largest], chas[i] = chas[i], chas[largest]
		} else {
			break
		}
		i = largest
		left = i*2 + 1
		right = i*2 + 2
	}
}

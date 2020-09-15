package sort

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(size, value int) []int {
	if size < 1 {
		return nil
	}

	arr := make([]int, size)
	next := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		arr[i] = next.Intn(value+1) - next.Intn(value)
	}
	return arr
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func GetMax(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}
	mid := left + (right-left)>>1
	maxLeft := GetMax(arr, left, mid)
	maxRight := GetMax(arr, mid+1, right)
	return max(maxRight, maxLeft)
}

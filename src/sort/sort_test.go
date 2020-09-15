package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBubblingSort(t *testing.T) {
	//arr := []int{1, 2, 9, 3, 5, 7, 11, 3, 3, 19, 8}
	//fmt.Println("before: ", arr)
	//BubblingSort(arr)
	//fmt.Println("after: ", arr)

	arr := GenerateRandomArray(20, 300)
	for i := 0; i < 20; i++ {
		seed := (i + 1) * 100
		rands := rand.New(rand.NewSource(time.Now().UnixNano()))
		arr[i] = rands.Intn(seed) - rands.Intn(seed-1)
	}
	fmt.Println("before: ", arr)
	BubblingSort(arr)
	fmt.Println("after: ", arr)
}

func TestSelectSort(t *testing.T) {
	//arr := []int{1, 2, 9, 3, 5, 7, 11, 3, 3, 19, 8}
	//fmt.Println("before: ", arr)
	//BubblingSort(arr)
	//fmt.Println("after: ", arr)

	arr := GenerateRandomArray(20, 300)
	fmt.Println("before: ", arr)
	SelectSort(arr)
	fmt.Println("after: ", arr)
}

func TestInsertSort(t *testing.T) {
	arr := GenerateRandomArray(20, 300)
	fmt.Println("before: ", arr)
	InsertSort(arr)
	fmt.Println("after: ", arr)
}
func TestMergeSort(t *testing.T) {
	arr := GenerateRandomArray(20, 300)
	fmt.Println("before: ", arr)
	MergeSort(arr)
	fmt.Println("after: ", arr)
}
func TestMergeSortDemo(t *testing.T) {
	DefaultTest(MergeSort)
}

func DefaultTest(handle func(arr []int)) {
	arr := GenerateRandomArray(20, 300)
	fmt.Println("before: ", arr)
	handle(arr)
	fmt.Println("after: ", arr)

}

func TestQuickSort(t *testing.T) {
	DefaultTest(QuickSort)

}

func TestHeapSort(t *testing.T) {
	DefaultTest(HeapSort)
}

func TestName(t *testing.T) {
	demo()
}

func TestDemo2(t *testing.T) {
	demo4()
}

func TestDemo5(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before: ", arr)
	demo5(&arr)
	fmt.Println("After: ", arr)
}

package sort

import "fmt"

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) {
	if left < right {
		start, end := partition(arr, left, right)
		start--
		end++
		quickSort(arr, left, start)
		quickSort(arr, end, right)
	}
}

func partition(arr []int, left, right int) (int, int) {
	less := left - 1
	more := right
	for left < more {
		if arr[left] < arr[right] {
			less++
			arr[less], arr[left] = arr[left], arr[less]
			left++
		} else if arr[left] > arr[right] {
			more--
			arr[left], arr[more] = arr[more], arr[left]
		} else {
			left++
		}
	}
	arr[more], arr[right] = arr[right], arr[more]
	less++
	return less, more
}

func demo() {
	var a int
	var b int
	a = 5
	b = 3
	fmt.Println(a / b)
}

func demo2() {
	var a int
	var b int
	a = 5
	b = 3
	var c = float64(a / b)
	fmt.Println(c)

}

func demo3(handle func(a, b int) int, a, b int) int {
	return handle(a, b)
}

func demo4() {
	i := demo3(func(a, b int) int { return a * b }, 1, 2)
	fmt.Println(i)
}

func demo5(a *[5]int) {
	a[3] = 0
}

func demo6(a []int) {
	if len(a) < 4 {
		panic("error")
	}
	a[3] = 0
}

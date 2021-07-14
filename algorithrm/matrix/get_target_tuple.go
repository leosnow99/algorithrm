package matrix

import "fmt"

/**
不重复打印排序数组中相加和为给定值的所有二元组和三
元组
【题目】
给定排序数组arr 和整数k，不重复打印arr 中所有相加和为k 的不降序二元组。
例如，arr=[-8,-4,-3,0,1,2,4,5,8,9]，k=10，打印结果为：
	1,9
	2,8

补充问题：给定排序数组arr 和整数k，不重复打印arr 中所有相加和为k 的不降序三元组。
例如，arr=[-8,-4,-3,0,1,2,4,5,8,9]，k=10，打印结果为：
	-4,5,9
	-3,4,9
	-3,5,8
	0,1,9
	0,2,8
	1,4,5
*/

func printUniquePair(arr []int, target int) {
	if len(arr) == 0 {
		return
	}

	left, right := 0, len(arr)-1
	for left < right {
		if arr[left]+arr[right] < target {
			left++
		} else if arr[left]+arr[right] > target {
			right--
		} else {
			if left == 0 || arr[left] != arr[left-1] {
				fmt.Println(arr[left], ",", arr[right])
			}
			left++
			right--
		}
	}
}

func printUniqueTriad(arr []int, target int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr)-2; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			printResult(arr, i, i+1, len(arr)-1, target-arr[i])
		}
	}
}

func printResult(arr []int, f, l, r, k int) {
	for l < r {
		if arr[l]+arr[r] < k {
			l++
		} else if arr[l]+arr[r] > k {
			r--
		} else {
			if (l == r+1) || arr[l] == arr[l-1] {
				fmt.Println(arr[l], ",", arr[r])
			}
			l++
			r--
		}
	}
}

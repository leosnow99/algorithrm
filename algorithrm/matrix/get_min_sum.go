package matrix

/**
计算数组的小和
【题目】
数组小和的定义如下：
	例如，数组s=[1,3,5,2,4,6]，在s[0]的左边小于或等于s[0]的数的和为0；
	在s[1]的左边小于或等于s[1]的数的和为1；在s[2]的左边小于或等于s[2]的数的和为1+3=4；
	在s[3]的左边小于或等于s[3]的数的和为1；在s[4]的左边小于或等于s[4]的数的和为1+3+2=6；
	在s[5]的左边小于或等于s[5]的数的和为1+3+5+2+4=15。所以s 的小和为0+1+4+1+6+15=27。
*/

// 下面介绍一种时间复杂度为O(NlogN)、额外空间复杂度为O(N)的方法:
// 	这是一种在归并排序的过程中，利用组间在进行合并时产生小和的过程。
//		1．假设左组为l[]，右组为r[]，左右两个组的组内都已经有序，现在要利用外排序合并成一个大组，并假设当前外排序是l[i]与r[j]在进行比较。
//		2．如果l[i]<=r[j]，那么产生小和。假设从r[j]往右一直到r[]结束，元素的个数为m，那么产生的小和为l[i]*m。
//		3．如果l[i]>r[j]，不产生任何小和。
//		4．整个归并排序的过程该怎么进行就怎么进行，排序过程没有任何变化，只是利用步骤1~步骤3，也就是在组间合并的过程中累加所有产生的小和，
//	      总的累加和就是结果。
func getSmallSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return process(arr, 0, len(arr) - 1)
}

// 在归并排序中，尤其是在组与组之间进行外排序合并的过程中，按照如上方式把小和一点一点地“榨”出来，最后收集到所有的小和
func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := (l + r) / 2
	return process(arr, l, mid) + process(arr, mid+1, r) + merge(arr, l, mid, r)
}

func merge(arr []int, left, mid, right int) int {
	h := make([]int, right-left+1)
	hi, i, j, smallSum := 0, left, mid, 0

	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			smallSum += arr[i] * (right - j + 1)
			h[hi] = arr[i]
			hi++
			i++
		} else {
			h[hi] = arr[j]
			hi++
			j++
		}
	}

	for j < (right+1) || i < (mid+1) {
		if i > mid {
			h[hi] = arr[j]
		} else {
			h[hi] = arr[i]
		}
		hi++
		i++
		j++
	}

	// 回写原数组
	for k := 0; k != len(h); k++ {
		arr[left] = h[k]
		left++
	}

	return smallSum
}

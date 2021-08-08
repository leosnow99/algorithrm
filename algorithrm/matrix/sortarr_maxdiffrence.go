package matrix

import (
	"algorithm/util"
	"math"
)

/**
数组排序之后相邻数的最大差值
【题目】
给定一个整型数组arr，返回排序后相邻两数的最大差值。
【举例】
arr=[9,3,1,10]。如果排序，结果为[1,3,9,10]，9 和3 的差为最大差值，故返回6。
arr=[5,5,5,5]。返回0。
【要求】
如果 arr 的长度为N，请做到时间复杂度为O(N)。
*/

/**
本题如果用排序法实现，其时间复杂度是O(NlogN)，而如果利用桶排序的思想（不是直接进行桶排序），可以做到时间复杂度为O(N)，额外空间复杂度为O(N)。
遍历arr 找到最小值和最大值，分别记为min 和max。如果arr 的长度为N，那么我们准备N+1 个桶，把max 单独放在第N+1 号桶里。
arr 中在[min,max)范围上的数放在1~N 号桶里，对于1~N 号桶中的每一个桶来说，负责的区间大小为(max-min)/N。
比如长度为10 的数组arr 中，最小值为10，最大值为110。那么就准备11 个桶，arr 中等于110 的数全部放在第11 号桶里。
区间[10,20)的数全部放在1 号桶里，区间[20,30)的数全部放在2 号桶里……，区间[100,110)的数全部放在10 号桶里。那么如果一个数为num，
它应该分配进(num - min) × len / (max - min)号桶里。

arr 一共有N 个数，min 一定会放进1 号桶里，max 一定会放进最后的桶里。所以，如果把所有的数放入N+1 个桶中，必然有桶是空的。
如果arr 经过排序，相邻的数有可能此时在同一个桶中，也可能在不同的桶中。在同一个桶中的任何两个数的差值都不会大于区间值，
而在空桶左右两边不空的桶里，相邻数的差值肯定大于区间值。所以产生最大差值的两个相邻数肯定来自不同的桶。所以只要计算桶之间数的间距就可以，
也就是只用记录每个桶的最大值和最小值，最大差值只可能来自某个非空桶的最小值减去前一个非空桶的最大值。
*/
func maxGap(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	length := len(arr)
	min, max := math.MinInt32, math.MaxInt32
	for _, n := range arr {
		min = util.Min(min, n)
		max = util.Max(max, n)
	}
	if min == max {
		return 0
	}

	hasNum := make([]bool, length+1)
	maxList := make([]int, length+1)
	minList := make([]int, length+1)
	bid := 0

	for i := 0; i < length; i++ {
		bid = bucket(int64(arr[i]), int64(length), int64(min), int64(max))
		if hasNum[bid] {
			minList[bid] = util.Min(minList[bid], arr[bid])
			maxList[bid] = util.Max(maxList[bid], arr[bid])
		} else {
			min = arr[bid]
			max = arr[bid]
		}

		hasNum[bid] = true
	}

	res, lastMax, i := 0, maxList[0], 1
	for ; i < length; i++ {
		if hasNum[i] {
			res = util.Max(res, minList[i]-lastMax)
			lastMax = maxList[i]
		}
	}

	return res
}

func bucket(num, len, max, min int64) int {
	return int((num - min) * len / (max - min))
}

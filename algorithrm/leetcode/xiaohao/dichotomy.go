package xiaohao

import (
	"sort"
)

//leet:875 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, H int) int {
	if H == 1 {
		return piles[0]
	}

	left := 1
	right := 1
	for _, value := range piles {
		if value > right {
			right = value
		}
	}

	for left < right {
		mid := (left + right) >> 1
		if canEat(piles, H, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canEat(piles []int, H int, speed int) bool {
	needTime := 0

	for _, value := range piles {
		needTime += (value + speed - 1) / speed
	}
	return needTime <= H
}

// leet:69 x的平方根
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 2 {
		return 1
	}
	left := 1
	right := x / 2
	for left < right {
		mid := (left+right)/2 + 1
		if (x / mid) == mid {
			return mid
		}
		if (x / mid) > mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

//leet:278 第一个错误的版本
func firstBadVersion(n int) int {
	if n == 1 {
		if isBadVersion(n) {
			return 0
		}
		return 1
	}
	left := 0
	right := n
	for left <= right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) {
			left = mid + 1
		} else {
			return left
		}
	}
	return 0
}

func isBadVersion(version int) bool {
	return true
}

//leet:475 供暖器
func findRadius(houses []int, heaters []int) int {
	res := 0
	size := len(heaters)
	sort.Ints(heaters)
	for _, value := range houses {
		//二分查找不小于value的第一个值
		left := 0
		right := size
		for left < right {
			mid := left + (right-left)>>1
			if value > heaters[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}
		lastLength := 1<<31 - 1
		nextLength := 1<<31 - 1
		if right != 0 {
			lastLength = abs(value - heaters[right-1])
		}
		if right != size {
			nextLength = abs(value - heaters[right])
		}
		lastLength = min(nextLength, lastLength)
		res = max(lastLength, res)

	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

//leet:704 二分查找
func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}

//153 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

//154 寻找旋转排序数组中的最小值 II
func findMin2(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return left
}

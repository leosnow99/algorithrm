package xiaohao

import "sort"

//leet:15 三数之和
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	result := make([][]int, 0)

	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}

	for i := 0; i < len(nums)-2; i++ {
		var left, right int
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		left = i + 1
		right = len(nums) - 1
		target := 0 - nums[i]
		for left < right {
			tem := nums[left] + nums[right]
			if tem < target {
				left++
			}
			if tem > target {
				right--
			}
			if tem == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

package simple

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	if length < 2 {
		return
	}
	k = k % length
	copy(nums, append(nums[length-k:], nums[:length-k]...))
}

func containsDuplicate(nums []int) bool {
	contain := make(map[int]bool)
	for _, num := range nums {
		if contain[num] == false {
			contain[num] = true
		} else {
			return true
		}
	}
	return false
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func intersect(nums1 []int, nums2 []int) []int {
	tem := make(map[int]int)
	for _, num := range nums1 {
		tem[num]++
	}
	result := make([]int, 0)
	for _, num := range nums2 {
		if tem[num] > 0 {
			tem[num]--
			result = append(result, num)
		}
	}
	return result
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			continue
		} else {
			return digits
		}
	}
	return append([]int{1}, digits...)
}

func plusOneDemo2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			continue
		} else {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}
	for i := 0; i < len(nums)-1; i++ {
		need := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if need == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumDemo2(nums []int, target int) []int {
	tem := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		need := target - nums[i]
		if tem[need] > 0 {
			return []int{i, tem[need]}
		}
		tem[nums[i]] = i
	}
	return nil
}

func isValidSudoku(board [][]byte) bool {
	var column = make([]map[byte]bool, 9)
	var row = make([]map[byte]bool, 9)
	var box = make([]map[byte]bool, 9)
	for rowIndex, rowArray := range board {
		row[rowIndex] = make(map[byte]bool)
		for columnIndex, value := range rowArray {
			if value == '.' {
				continue
			}
			if column[columnIndex] == nil {
				column[columnIndex] = make(map[byte]bool)
			}
			_, cOk := column[columnIndex][value]
			if !cOk {
				column[columnIndex][value] = true
			} else {
				fmt.Println("column")
				return false
			}
			_, rOk := row[rowIndex][value]
			if !rOk {
				row[rowIndex][value] = true
			} else {
				return false
			}
			boxIndex := (rowIndex/3)*3 + (columnIndex / 3)
			if box[boxIndex] == nil {
				box[boxIndex] = make(map[byte]bool)
			}
			_, bOk := box[boxIndex][value]
			if !bOk {
				box[boxIndex][value] = true
			} else {
				return false
			}
		}
	}
	return true
}

func isValidSudokuDemo2(board [][]byte) bool {
	var column = make([]map[byte]bool, 9)
	var row = make([]map[byte]bool, 9)
	var box = make([]map[byte]bool, 9)

	for rowIndex, rowArray := range board {
		row[rowIndex] = make(map[byte]bool)
		for columnIndex, value := range rowArray {
			if board[rowIndex][columnIndex] != '.' {
				if row[rowIndex][value] {
					return false
				} else {
					row[rowIndex][value] = true
				}

				if column[columnIndex] == nil {
					column[columnIndex] = make(map[byte]bool)
				}
				if column[columnIndex][value] {
					return false
				} else {
					column[columnIndex][value] = true
				}

				boxIndex := (columnIndex / 3) + (rowIndex/3)*3
				if box[boxIndex] == nil {
					box[boxIndex] = make(map[byte]bool)
				}
				if box[boxIndex][value] {
					return false
				} else {
					box[boxIndex][value] = true
				}
			}
		}
	}
	return true
}

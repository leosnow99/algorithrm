package string_type

import "strings"

// 在有序但含有空的数组中查找字符串

// 1．假设在strs[left..right]上进行查找的过程，全局整型变量res 表示字符串str 在strs 中最
//左的位置。初始时，left=0，right=strs.length-1，res=-1。
//2．令mid=(left+right)/2，则strs[mid]为strs[left..right]中间位置的字符串。
//3．如果字符串strs[mid]与str 一样，说明找到了str，令res=mid。但要找的是最左的位置，
//还要在左半区寻找，看有没有更左的str 出现，所以令right=mid-1，然后重复步骤2。
//4．如果字符串strs[mid]与str 不一样，并且strs[mid]!=null，此时可以比较strs[mid]和str，
//如果strs[mid]的字典顺序比str 小，说明整个左半区不会出现str，需要在右半区寻找，所以令
//left=mid+1，然后重复步骤2。
// 5．如果字符串strs[mid]与str 不一样，并且strs[mid]==null，此时从mid 开始，从右到左遍
//历左半区（即strs[left..mid]）。如果整个左半区都为null，那么继续用二分的方式在右半区上查
//找（即令left=mid+1），然后重复步骤2。如果整个左半区不都为null，假设从右到左遍历
//strs[left..mid]时，发现第一个不为null 的位置是i，那么把str 和strs[i]进行比较。如果strs[i]字典
//顺序小于str，同样说明整个左半区没有str，令left=mid+1，然后重复步骤2。如果strs[i]字典顺
//序等于str，说明找到str，令res=mid，但要找的是最左的位置，还要在strs[left..i-1]上寻找，看
//有没有更左的str 出现，所以令right=i-1，然后重复步骤2。如果strs[i]字典顺序大于str，说明
//strs[i..right]上都没有str，需要在strs[left..i-1]上查找，所以令right=i-1，然后重复步骤2。
func getIndex(strs []string, str string) int {
	if len(strs) == 0 || len(str) == 0 {
		return -1
	}
	res := -1
	left, right := 0, len(strs)-1
	mid := 0
	i := 0
	for left <= right {
		mid = (left + right) >> 1
		if len(strs) != 0 && strs[mid] == str {
			res = mid
			right = mid - 1
		} else if len(strs[mid]) != 0 {
			if strings.Compare(strs[mid], str) < 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			i = mid
			i--
			for len(strs[i]) == 0 && i >= left {
				i--
			}
			if i < left || strings.Compare(strs[i], str) < 0 {
				left = mid + 1
			} else {
				if strs[i] == str {
					res = i
				}
				right = i - 1
			}
		}
	}
	return res
}

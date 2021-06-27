package string_type

import (
	"algorithm/util"
	"math"
)

// 数组中两个字符串的最小距离

// 给定一个字符串数组strs，再给定两个字符串str1 和str2，返回在strs 中str1 与str2 的最
// 小距离，如果str1 或str2 为null，或不在strs 中，返回-1。
func miniDistance(strs []string, str1, str2 string) int {
	if len(strs) == 0 || len(str1) == 0 || len(str2) == 0 {
		return -1
	}
	lastStr1, lastStr2 := -1, -1
	min := math.MaxInt32

	for i := 0; i < len(strs); i++ {
		if strs[i] == str1 {
			if lastStr2 != -1 {
				min = util.Min(min, i-lastStr2)
			}
			lastStr1 = i
		}
		if strs[i] == str2 {
			if lastStr1 != -1 {
				min = util.Min(min, i-lastStr1)
			}
			lastStr2 = i
		}
	}
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

// 进阶问题：如果查询发生的次数有很多，如何把每次查询的时间复杂度降为O(1)

// 其实是通过数组strs 先生成某种记录，在查询时通过记录进行查询。本书提供
// 了一种记录的结构供读者参考，如果strs 的长度为N，那么生成记录的时间复杂度为O(N2)，记
// 录的空间复杂度为O(N2)，

type Record struct {
	record map[string]map[string]int
}

func (r Record) update(indexMap map[string]int, str string, i int) {
	if _, ok := r.record[str]; !ok {
		r.record[str] = make(map[string]int)
	}
	strMap := r.record[str]
	for key, index := range indexMap {
		if key != str {
			lastMap := r.record[key]
			curMin := i - index
			if preMin, ok := strMap[key]; ok {
				if curMin < preMin {
					strMap[key] = curMin
					lastMap[str] = curMin
				}
			} else {
				strMap[key] = curMin
				lastMap[str] = curMin
			}
		}
	}
}

func NewRecord(strArr []string) Record {
	r := Record{}
	r.record = make(map[string]map[string]int)
	indexMap := make(map[string]int)
	for i := 0; i < len(strArr); i++ {
		curStr := strArr[i]
		r.update(indexMap, curStr, i)
		indexMap[curStr] = i
	}
	return r
}

func (r Record) miniDistance(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		return -1
	}
	if str1 == str2 {
		return 0
	}
	if v, ok := r.record[str1]; ok {
		if mini, ok := v[str2]; ok {
			return mini
		}
	}
	return -1
}

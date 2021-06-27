package string_type

import (
	"sort"
	"strings"
)

/**
拼接所有字符串产生字典顺序最小的大写字符串

给定一个字符串类型的数组strs，请找到一种拼接顺序，使得将所有的字符串拼接起来组
成的大写字符串是所有可能性中字典顺序最小的，并返回这个大写字符串。
*/

type dictString []string

func (d dictString) Len() int {
	return len(d)
}

func (d dictString) Less(i, j int) bool {
	return d[i]+d[j] < d[j]+d[i]
}

func (d dictString) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func lowerString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Sort(dictString(strs))
	var res strings.Builder
	for _, s := range strs {
		res.WriteString(s)
	}
	return res.String()
}

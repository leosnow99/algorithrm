package string_type

// 判断两个字符串是否互为变形词
//
// 给定两个字符串str1 和str2，如果str1 和str2 中出现的字符种类一样且每种字符出现的次
// 数也一样，那么str1 与str2 互为变形词。请实现函数判断两个字符串是否互为变形词。

func isDeformation(str1, str2 string) bool {
	if len(str1) == 0 || len(str2) == 0 || len(str1) != len(str2) {
		return false
	}
	ch1 := []byte(str1)
	ch2 := []byte(str2)
	var maps = make([]uint8, 256)
	for _, c := range ch1 {
		maps[c]++
	}
	for _, c2 := range ch2 {
		if maps[c2] == 0 {
			return false
		} else {
			maps[c2]--
		}
	}
	return true
}

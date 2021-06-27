package string_type

import "strconv"

// 字符串的统计字符串

// 给定一个字符串 str，返回str 的统计字符串。例如，"aaabbadddffc"的统计字符串为
//"a_3_b_2_a_1_d_3_f_2_c_1"。

func getCountString(str string) string {
	if len(str) == 0 {
		return ""
	}
	chs := []byte(str)
	res := string(chs[0])
	num := 1
	for i := 1; i < len(chs); i++ {
		if chs[i] != chs[i-1] {
			res = concat(res, strconv.Itoa(num), string(chs[i]))
		} else {
			num++
		}
	}
	return concat(res, strconv.Itoa(num), "")
}

func concat(str1, str2, str3 string) string {
	if str3 == "" {
		return str1 + "_" + str2
	}
	return str1 + "_" + str2 + "_" + str3
}

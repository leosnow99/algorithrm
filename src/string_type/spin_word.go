package string_type

import "strings"

// 判断两个字符串是否互为旋转词

// 如果a 和b 的长度不一样，字符串a 和b 不可能互为旋转词。如果
// a 和b 长度一样，先生成一个大字符串b2，b2 是两个字符串b 拼在一起的结果，即String b2 = b + b。
// 然后看b2 中是否包含字符串a，如果包含，说明字符串a 和b 互为旋转词，否则说明两个 字符串不互为旋转词

func isRotation(str1, str2 string) bool {
	if len(str1) ==0 || len(str2)==0 || len(str1) != len(str2) {
		return false
	}
	a := str1 + str1
	return strings.Contains(a, str2)
}
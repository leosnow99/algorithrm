package string_type

// 翻转字符串
// 给定一个字符类型的数组chas，请在单词间做逆序调整。只要做到单词的顺序逆序即可，
//对空格的位置没有特别要求。

func reverseWord(chas []byte) {
	if len(chas) == 0 {
		return
	}
	reverse(chas, 0, len(chas)-1)
	l, r := -1, -1
	for i := 0; i < len(chas); i++ {
		if chas[i] != ' ' {
			if i == 0 || chas[i-1] == ' ' {
				l = i
			}
			if i == len(chas)-1 || chas[i+1] == ' ' {
				r = i
			}
		}
		if l != -1 && r != -1 {
			reverse(chas, l, r)
			l = -1
			r = -1
		}
	}
}

func reverse(chas []byte, start, end int) {
	for start < end {
		chas[start], chas[end] = chas[end], chas[start]
		start++
		end--
	}
}

// 给定一个字符类型的数组chas 和一个整数size，请把大小为size 的左半区整体
//移到右半区，右半区整体移到左边。
func rotate(chas []byte, size int) {
	if len(chas) == 0 || size <= 0 || size >= len(chas) {
		return
	}
	reverse(chas, 0, size-1)
	reverse(chas, size, len(chas)-1)
	reverse(chas, 0, len(chas)-1)
}
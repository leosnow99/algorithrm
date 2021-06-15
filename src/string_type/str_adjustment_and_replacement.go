package string_type

//字符串的调整与替换  --- 逆序复制
//给定一个字符类型的数组chas[]，chas 右半区全是空字符，左半区不含有空字符。现在想
//将左半区中所有的空格字符替换成"%20"，假设chas 右半区足够大，可以满足替换所需要的空
//间，请完成替换函数。

func replace(chas []byte) {
	if len(chas) == 0 {
		return
	}
	num, lens := 0, 0
	for ; lens < len(chas) && chas[lens] != 0; lens++ {
		if chas[lens] == ' ' {
			num++
		}
	}
	j := lens + num*2 - 1
	for i := lens - 1; i > -1; i-- {
		if chas[i] != ' ' {
			chas[j] = chas[i]
			j--
		} else {
			chas[j] = '0'
			j--
			chas[j] = '2'
			j--
			chas[j] = '%'
			j--
		}
	}
}

// 给定一个字符类型的数组chas[]，其中只含有数字字符和“*”字符。现在想把
//所有的“*”字符挪到chas 的左边，数字字符挪到chas 的右边。
func modify(chas []byte) {
	if len(chas) == 0 {
		return
	}
	j := len(chas) - 1
	for i := len(chas) - 1; i > -1; i-- {
		if chas[i] != '*' {
			chas[j] = chas[i]
			j--
		}
	}
	for j > -1 {
		chas[j] = '*'
		j--
	}
}

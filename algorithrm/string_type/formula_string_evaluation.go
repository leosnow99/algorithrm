package string_type

import (
	"strconv"

	"github.com/ahrtr/gocontainer/stack"
)

// 公式字符串求值
// 给定一个字符串str，str 表示一个公式，公式里可能有整数、加减乘除符号和左右括号, 返回公式的计算结果。
// 说明：
// 1．可以认为给定的字符串一定是正确的公式，即不需要对str 做公式有效性检查。
// 2．如果是负数，就需要用括号括起来，比如"4*(-3)"。但如果负数作为公式的开头或括号部分的开头，则可以没有括号，比如"-3*4"和"(-3*4)"都是合法的。
// 3．不用考虑计算过程中会发生溢出的情况。

func getValue(str string) int {
	if len(str) == 0 {
		return 0
	}
	v, _ := value([]byte(str), 0)
	return v
}

func value(chars []byte, i int) (int, int) {
	dep := stack.New()
	pre := 0
	for i < len(chars) && chars[i] != ')' {
		if chars[i] >= '0' && chars[i] <= '9' {
			pre = pre*10 + int(chars[i]-'0')
			i++
		} else if chars[i] != '(' {
			addNum(dep, pre)
			dep.Push(chars[i])
			i++
			pre = 0
		} else {
			pre, i = value(chars, i+1)
		}
	}
	addNum(dep, pre)
	return getNum(dep), i
}

func addNum(deq stack.Interface, num int) {
	if !deq.IsEmpty() {
		cur := 0
		top := deq.Pop().(string)
		if top == "+" || top == "-" {
			deq.Push(top)
		} else {
			cur, _ = strconv.Atoi(top)
			if top == "*" {
				num *= cur
			} else {
				num /= cur
			}
		}
	}
	deq.Push(strconv.Itoa(num))
}

func getNum(deq stack.Interface) int {
	var res = 0
	var cur string
	add := true
	for !deq.IsEmpty() {
		cur = deq.Pop().(string)
		if cur == "+" {
			add = true
		} else if cur == "-" {
			add = false
		} else {
			num, _ := strconv.Atoi(cur)
			if add {
				res += num
			} else {
				res -= num
			}
		}
	}
	return res
}

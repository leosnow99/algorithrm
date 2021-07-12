package matrix

import "fmt"

/**
在数组中找到出现次数大于N/K 的数
【题目】
给定一个整型数组arr，打印其中出现次数大于一半的数，如果没有这样的数，打印提示信息。
进阶问题：给定一个整型数组arr，再给定一个整数K，打印所有出现次数大于N/K 的数，
如果没有这样的数，打印提示信息。
*/

// 第一个for 循环就是一次在数组中删掉两个不同数的代码实现。我们把变量 cand 叫作候选，times 叫作次数
// 	 times==0 时，表示当前没有候选，则把当前数arr[i]设成候选，同时把times 设置成1。
// 	 times!=0 时，表示当前有候选，如果当前的数arr[i]与候选一样，就把times 加1；如
// 果当前的数arr[i]与候选不一样，就把times 减1，减到0 则表示又没有候选了
func printHalfMajor(arr []int) int {
	cand, times := 0, 0
	// 第一个for 循环的实质就是我们的核心解题思路，一次在数组中删掉两个不同的数，不停地删除，
	// 直到剩下的数只有一种，如果一个数出现次数大于一半，则这个数最后一定会被剩下来，也就是最后的cand 值。
	for i := 0; i != len(arr); i++ {
		if times == 0 {
			cand = arr[i]
			times++
		} else if cand == arr[i] {
			times++
		} else {
			times--
		}
	}

	times = 0
	for i := 0; i != len(arr); i++ {
		if cand == arr[i] {
			times++
		}
	}
	if times >= len(arr)/2 {
		return cand
	}

	fmt.Println("no such number")
	return 0
}

func printKMajor(arr []int, k int) []int {
	if k < 2 {
		fmt.Println("the value of K is invalid.")
		return nil
	}

	maps := make(map[int]int)
	for _, number := range arr {
		if v, ok := maps[number]; ok {
			maps[number] = v + 1
		} else if len(maps) == k-1 {
			allCandsMinusOne(maps)
		} else {
			maps[number] = 1
		}
	}

	var res []int
	reals := getReals(arr, maps)
	for number, v := range reals {
		if v > len(arr) / k {
			res = append(res, number)
		}
	}

	return res
}

func allCandsMinusOne(maps map[int]int) {
	for k, v := range maps {
		if v == 1 {
			delete(maps, k)
		}
		maps[k] = v - 1
	}
}

func getReals(arr []int, cands map[int]int) map[int]int {
	res := make(map[int]int)
	for _, number := range arr {
		if _, ok := cands[number]; ok {
			if curReal, cOk := res[number]; cOk {
				res[number] = curReal + 1
			} else {
				res[number] = 1
			}
		}
	}
	return res
}

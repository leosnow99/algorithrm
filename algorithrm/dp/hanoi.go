package dp

import "fmt"

func move(n int, from, mid, to string) {
	if n == 1 {
		fmt.Println("move from " + from + " to " + to)
	} else {
		move(n-1, from, to, mid)
		move(1, from, mid, to)
		move(n-1, mid, from, to)
	}
}

func hanoi(n int) {
	if 0 < n {
		move(n, "from", "mid", "to")
	}
}

/**
如果arr 长度为N，请实现时间复杂度为O(N)、额外空间复杂度为O(1)的方法。

首先求都在from 柱子上的圆盘1~i，如果都移动到to 上的最少步骤数假设为
S(i)。根据上面的步骤，S(i)=步骤1 的步骤总数+1+步骤3 的步骤总数=S(i-1)+1+S(i-1)，S(1)=1。
所以S(i)+1=2(S(i-1)+1)，S(1)+1==2。根据等比数列求和公式得到S(i)+1=2i，所以S(i)=2i-1。
*/
func step(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return hanoiProcess(arr, len(arr)-1, 1, 2, 3)
}

func hanoiProcess(arr []int, length, from, mid, to int) int {
	if length == -1 {
		return 0
	}

	if arr[length] != from && arr[length] != to {
		return -1
	}

	if arr[length] == from {
		return hanoiProcess(arr, length-1, from, to, mid)
	}

	rest := hanoiProcess(arr, length-1, mid, from, to)
	if rest == -1 {
		return rest
	}

	return (1 << rest) + rest
}

func step2(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	res, tmp := 0, 0
	from, mid, to := 1, 2, 3
	for i := len(arr) - 1; i >= 0; {
		if arr[i] != from && arr[i] != to {
			return -1
		}
		if arr[i] == to {
			res += 1 << i
			tmp = from
			from = mid
		} else {
			tmp = to
			to = mid
		}
		mid = tmp
		i--
	}
	return res
}

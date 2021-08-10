package matrix

import (
	"algorithm/datastruct/heaps"
	"container/heap"
)

/**
分金条的最小花费
【题目】
给定一个正数数组arr，arr 的累加和代表金条的总长度，arr 的每个数代表金条要分成的长
度。规定长度为K 的金条只需分成两块，费用为K 个铜板。返回把金条分出arr 中的每个数字
需要的最小代价。
【举例】
arr={10,30,20}，金条总长度为60。
如果先分成40 和20 两块，将花费60 个铜板，再把长度为40 的金条分成10 和30 两块，
将花费40 个铜板，总花费为100 个铜板；如果先分成10 和50 两块，将花费60 个铜板，再把
长度为50 的金条分成20 和30 两块，将花费50 个铜板，总花费为110 个铜板；如果先分成30
和30 两块，将花费60 个铜板，再把其中一根长度为30 的金条分成10 和20 两块，将花费30
个铜板，总花费为90 个铜板。所以返回最低花费为90。
*/

/**
原型为哈夫曼编码算法，是用贪心策略求解的
0．假设最小代价为ans，初始时ans=0。先把arr 中所有的数字放进一个小根堆。
1．从小根堆中弹出两个数字，假设为a 和b，令ans=ans+a+b，然后把a+b 的和放进小根堆。
2．重复步骤1，直到小根堆中只剩一个数字过程停止，返回ans 即可。
*/

func getMinSplitCost(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	minHeap := &heaps.IntHeap{
		Heap:    make([]int, len(arr)),
		MinRoot: true,
	}
	for i := 0; i < len(arr); i++ {
		heap.Push(minHeap, arr[i])
	}

	var ans int
	for minHeap.Len() != 0 {
		sum := minHeap.Pop().(int) + minHeap.Pop().(int)
		ans += sum
		heap.Push(minHeap, sum)
	}

	return ans
}

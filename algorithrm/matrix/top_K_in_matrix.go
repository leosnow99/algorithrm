package matrix

import "fmt"

/**
打印 N 个数组整体最大的Top K
【题目】
有 N 个长度不一的数组，所有的数组都是有序的，请从大到小打印这N 个数组整体最大的前K 个数。

例如，输入含有N 行元素的二维数组可以代表N 个一维数组。
	219,405,538,845,971
	148,558
	52,99,348,691
再输入整数k=5，则打印：
	Top 5: 971,845,691,558,538
*/

/**
本题的解法是利用堆结构和堆排序的过程完成的，具体过程如下：
	1．构建一个大小为N 的大根堆heap，建堆的过程就是把每个数组中的最后一个值（也就是该数组的最大值）依次加入堆里，
		这个过程是建堆时的调整过程（heapInsert）。
	2．建好堆之后，此时heap 堆顶的元素是所有数组的最大值中最大的那个，打印堆顶元素。
	3．假设堆顶元素来自a 数组的i 位置。那么接下来就把堆顶的前一个数（即a[i-1]）放在heap 的头部，
		也就是用a[i-1]替换原本的堆顶，然后从堆的头部开始调整堆，使其重新变为大根堆（heapify 过程）。
	4．这样每次都可以得到一个堆顶元素max，在打印完成后都经历步骤3 的调整过程。整体打印k 次，就是从大到小全部的Top K。
	5．在重复步骤3 的过程中，如果max 来自的那个数组（仍假设是a 数组）已经没有元素。也就是说，max 已经是a[0]，再往左没有数了。
		那么就把heap 中最后一个元素放在heap 头部的位置，然后把heap 的大小减1（heapSize-1），最后依然是从堆的头部开始调整堆，
		使其重新变为大根堆（堆大小减1 之后的heapify 过程）。
	6．直到打印了k 个数，过程结束。
*/

type heapNode struct {
	value  int // 值是什么
	arrNum int // 来自哪个数组
	index  int // 来自数组的哪个位置
}

func newHeapNode(value, arrNum, index int) heapNode {
	return heapNode{
		value:  value,
		arrNum: arrNum,
		index:  index,
	}
}

func printTopKInNArr(matrix [][]int, topK int) {
	heapSize := len(matrix)
	heap := make([]heapNode, heapSize)

	for i := 0; i < heapSize; i++ {
		index := len(matrix[i]) - 1
		heap[i] = newHeapNode(matrix[i][index], i, index)
		headHeapInsert(heap, i)
	}

	fmt.Printf("TOP %d : \n", topK)
	for i := 0; i < topK; i++ {
		if heapSize == 0 {
			break
		}

		fmt.Printf("%d ", heap[0])

		if heap[0].index != 0 {
			heap[0].index--
			heap[0].value = matrix[heap[0].arrNum][heap[0].index]
		} else {
			heapSize--
			heap[0], heap[heapSize] = heap[heapSize], heap[0]
		}

		headHeapify(heap, 0, heapSize)
	}
}

func headHeapInsert(heap []heapNode, index int) {
	for index != 0 {
		parent := (index - 1) / 2
		if heap[parent].value < heap[index].value {
			heap[parent], heap[index] = heap[index], heap[parent]
			index = parent
		} else {
			break
		}
	}
}

func headHeapify(heap []heapNode, index, heapSize int) {
	left, right := index>>1+1, index>>1+2
	largest := index

	for left < heapSize {
		if heap[index].value < heap[left].value {
			largest = left
		}

		if right < heapSize && heap[largest].value < heap[right].value {
			largest = right
		}

		if largest != index {
			heap[largest], heap[index] = heap[index], heap[largest]
		} else {
			break
		}

		index = largest
		left = index>>1 + 1
		right = index>>1 + 2
	}
}

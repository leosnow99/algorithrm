package matrix

import (
	"container/heap"
)

type program struct {
	cost   int
	profit int
}

func newProgram(cost, profit int) program {
	return program{
		cost:   cost,
		profit: profit,
	}
}

func getMaxMoney(W, K int, costs, profits []int) int {
	if W < 1 || K < 0 || len(costs) == 0 || len(profits) == 0 || len(costs) != len(profits) {
		return W
	}

	// 项目花费小根堆
	costMinHeap := newCostSmallRootHeap(len(costs))
	// 项目利润大根堆
	profitMaxHeap := newProfitMaxRootHeap(len(costs))

	for i := 0; i < len(costs); i++ {
		heap.Push(costMinHeap, newProgram(costs[i], profits[i]))
	}

	// 依次做 K 个项目
	for i := 1; i <= K; i++ {
		// 当前资金为W，在项目花费小根堆里所有花费小于或等于W 的项目，都可以考虑
		for costMinHeap.Len() != 0 && (*costMinHeap)[0].cost <= W {
			// 把可以考虑的项目都放进项目利润大根堆里
			heap.Push(profitMaxHeap, heap.Pop(costMinHeap).(program))
		}

		// 如果此时项目利润大根堆为空，说明可以考虑的项目为空
		// 说明当前资金W 已经无法解锁任何项目，直接返回W
		if profitMaxHeap.Len() == 0 {
			return W
		}

		W += heap.Pop(profitMaxHeap).(program).profit
	}

	return W
}

type costSmallRootHeap []program

func newCostSmallRootHeap(length int) *costSmallRootHeap {
	t := costSmallRootHeap(make([]program, length))
	return &t
}

func (c *costSmallRootHeap) Len() int {
	return len(*c)
}

func (c *costSmallRootHeap) Less(i, j int) bool {
	return (*c)[i].cost < (*c)[j].cost
}

func (c *costSmallRootHeap) Swap(i, j int) {
	(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
}

func (c *costSmallRootHeap) Push(x interface{}) {
	*c = append(*c, x.(program))
}

func (c *costSmallRootHeap) Pop() interface{} {
	oldV := (*c)[c.Len()-1]
	*c = (*c)[:c.Len()-1]

	return oldV
}

type profitMaxRootHeap []program

func newProfitMaxRootHeap(length int) *profitMaxRootHeap {
	t := profitMaxRootHeap(make([]program, length))
	return &t
}

func (p *profitMaxRootHeap) Len() int {
	return len(*p)
}

func (p *profitMaxRootHeap) Less(i, j int) bool {
	return (*p)[i].profit > (*p)[j].profit
}

func (p *profitMaxRootHeap) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *profitMaxRootHeap) Push(x interface{}) {
	*p = append(*p, x.(program))
}

func (p *profitMaxRootHeap) Pop() interface{} {
	oldV := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]

	return oldV
}

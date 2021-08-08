package matrix

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

	return 0
}

type costSmallRootHeap []program

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

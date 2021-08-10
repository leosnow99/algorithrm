package heap

type IntHeap struct {
	Heap []int
	// true是小根堆，false是大根堆
	MinRoot bool
}

func (h *IntHeap) Len() int { return len(h.Heap) }

func (h *IntHeap) Less(i, j int) bool {
	if h.MinRoot {
		return h.Heap[i] < h.Heap[j]
	} else {
		return h.Heap[i] > h.Heap[j]
	}

}

func (h *IntHeap) Swap(i, j int) { h.Heap[i], h.Heap[j] = h.Heap[j], h.Heap[i] }

func (h *IntHeap) Push(x interface{}) {
	h.Heap = append(h.Heap, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h
	n := len(old.Heap)
	x := old.Heap[n-1]
	h.Heap = old.Heap[0 : n-1]
	return x
}

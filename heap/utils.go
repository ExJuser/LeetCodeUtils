package heap

type IntMinHeap []int

func (h *IntMinHeap) Len() int {
	return len(*h)
}

func (h *IntMinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

type IntMaxHeap []int

func (h *IntMaxHeap) Len() int {
	return len(*h)
}

func (h *IntMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

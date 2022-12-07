package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/12/6 14:31
 */

//小根堆
type IntMinHeap []int

func (h *IntMinHeap) Less(i, j int) bool {
	// h[i] > h[j]为大根堆
	return (*h)[i] < (*h)[j]
}

func (h *IntMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMinHeap) Len() int {
	return len(*h)
}

func (h *IntMinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *IntMinHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

//大根堆
type IntMaxHeap []int

func (h *IntMaxHeap) Less(i, j int) bool {
	// h[i] > h[j]为大根堆
	return (*h)[i] < (*h)[j]
}

func (h *IntMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntMaxHeap) Len() int {
	return len(*h)
}

func (h *IntMaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *IntMaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntMaxHeap) Size() int {
	return len(*h)
}

func (h *IntMaxHeap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *IntMaxHeap) Peek() (v interface{}) {
	if h.IsEmpty() {
		return -1
	}
	return (*h)[0]
}

package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/12/6 14:31
 */

type IntHeap struct {
	nums []int
	//大根堆还是小跟堆，true为大根堆
	isMaxHeap bool
}

func NewMaxIntHeap() *IntHeap {
	return &IntHeap{
		nums:      []int{},
		isMaxHeap: true,
	}
}

func NewMinIntHeap() *IntHeap {
	return &IntHeap{
		nums:      []int{},
		isMaxHeap: false,
	}
}

func (h *IntHeap) Less(i, j int) bool {
	if h.isMaxHeap {
		// h[i] > h[j]为大根堆
		return h.nums[i] > h.nums[j]
	}
	return h.nums[i] < h.nums[j]
}

func (h *IntHeap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *IntHeap) Len() int {
	return len(h.nums)
}

func (h *IntHeap) Pop() interface{} {
	if h.IsEmpty() {
		return -1
	}
	top := h.nums[h.Len()-1]
	h.nums = h.nums[:h.Len()-1]
	return top
}

func (h *IntHeap) Push(v interface{}) {
	h.nums = append(h.nums, v.(int))
}

func (h *IntHeap) Size() int {
	return h.Len()
}

func (h *IntHeap) IsEmpty() bool {
	return h.Size() == 0
}

func (h *IntHeap) Peek() (v interface{}) {
	if h.IsEmpty() {
		return -1
	}
	return h.nums[0]
}

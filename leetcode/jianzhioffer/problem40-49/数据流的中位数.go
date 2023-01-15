package problem40_49

import (
	"container/heap"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/13 17:29
 */

//如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数
//值排序之后中间两个数的平均值。
//
// 例如，
//
// [2,3,4] 的中位数是 3
//
// [2,3] 的中位数是 (2 + 3) / 2 = 2.5
//
// 设计一个支持以下两种操作的数据结构：
//
//
// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
//
//
// 示例 1：
//
// 输入：
//["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
//[[],[1],[2],[],[3],[]]
//输出：[null,null,null,1.50000,null,2.00000]
//
//
// 示例 2：
//
// 输入：
//["MedianFinder","addNum","findMedian","addNum","findMedian"]
//[[],[2],[],[3],[]]
//输出：[null,null,2.00000,null,2.50000]
//
//
//
// 限制：
//
//
// 最多会对 addNum、findMedian 进行 50000 次调用。
//
//
// 注意：本题与主站 295 题相同：https://leetcode-cn.com/problems/find-median-from-data-
//stream/
//
// Related Topics 设计 双指针 数据流 排序 堆（优先队列） 👍 382 👎 0

/*
	思路：我们用一个小顶堆H1保存较大的一半数，一个大顶堆H2保存较小的一半数，且保证size(H1) >= size(H2)
		只需要通过size(H1) == size(H2)的时候，每次都往H1中加元素即可保证
	加入元素a：
		1.H1、H2均为空，H1.Push(a)；
		2.否则：
			我们先将元素a加到H2，再讲H2中最大的数加到H1中，即可保证H1和H2的关系
			H2.Push(a)
			H1.Push(H2.Pop())


*/
//leetcode submit region begin(Prohibit modification and deletion)
type IntHeap struct {
	nums []int
	//大根堆还是小跟堆，true为大根堆
	isMaxHeap bool
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
	return h.nums[len(h.nums)-1]
}

type MedianFinder struct {
	h1 *IntHeap
	h2 *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	h1, h2 := &IntHeap{nums: []int{}, isMaxHeap: false}, &IntHeap{nums: []int{}, isMaxHeap: true}
	heap.Init(h1)
	heap.Init(h2)
	return MedianFinder{
		h1: h1,
		h2: h2,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.h1.IsEmpty() && this.h2.IsEmpty() {
		this.h1.Push(num)
		return
	}
	this.h2.Push(num)
	this.h1.Push(this.h2.Pop())
}

func (this *MedianFinder) FindMedian() float64 {
	if this.h1.Size() != this.h2.Size() {
		return float64(this.h1.Peek().(int))
	}
	p1, p2 := this.h1.Peek().(int), this.h2.Peek().(int)
	return (float64(p1) + float64(p2)) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)

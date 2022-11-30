package LinkedList

import (
	"container/heap"
	"example.com/m/v2/tools/collections"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/30 11:16
 */

//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2258 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
跟合并两个一样，难点在如何找出k个值中的最小值，很容易想到最小堆
*/
func mergeKLists(lists []*collections.ListNode) *collections.ListNode {
	res := new(collections.ListNode)
	dummy := res
	minHeap := &ListNodeHeap{}

	for _, head := range lists {
		if head != nil {
			heap.Push(minHeap, head)
		}
	}

	heap.Init(minHeap)
	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(*collections.ListNode)
		dummy.Next = top
		if top.Next != nil {
			heap.Push(minHeap, top.Next)
		}
		dummy = dummy.Next
	}
	return res.Next
}

type ListNodeHeap []*collections.ListNode

func (h *ListNodeHeap) Less(i, j int) bool {
	// h[i] > h[j]为大根堆
	return (*h)[i].Val < (*h)[j].Val
}

func (h *ListNodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *ListNodeHeap) Len() int {
	return len(*h)
}

func (h *ListNodeHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *ListNodeHeap) Push(v interface{}) {
	*h = append(*h, v.(*collections.ListNode))
}

//leetcode submit region end(Prohibit modification and deletion)

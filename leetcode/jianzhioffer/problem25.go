package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/11/21 20:17
 */

//输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
//
// 输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
// 限制：
//
// 0 <= 链表长度 <= 1000
//
// 注意：本题与主站 21 题相同：https://leetcode-cn.com/problems/merge-two-sorted-lists/
//
// Related Topics 递归 链表 👍 295 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代做法，简单模拟即可
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	p, q, head := l1, l2, res
	for p != nil && q != nil {
		if p.Val <= q.Val {
			head.Next = p
			p = p.Next
		} else {
			head.Next = q
			q = q.Next
		}
		head = head.Next
	}

	if p != nil {
		head.Next = p
	}
	if q != nil {
		head.Next = q
	}
	return res.Next
}

//递归做法
/*
首先递归的终止条件，就是l1为空，返回l2；l2为空，返回l1。非常漂亮的逻辑

然后就很简单了，l1的后续结点不就是l1.Next和l2的合并吗？直接递归！返回l1，完事！,反之，换成l2即可
*/
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoLists2(l1.Next, l2)
	return l1
}

package LinkedList

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/2 11:03
 */
//给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
//
//提示：
//
//
// 链表中的节点数目为 n
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
//
//
//
//
// 进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
//
//
//
//
// Related Topics 递归 链表 👍 1857 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 详解参考： https://labuladong.github.io/algo/2/19/20/
*/
func reverseKGroup(head *collections.ListNode, k int) *collections.ListNode {
	a, b := head, head
	cnt := 0
	for b != nil && cnt < k {
		cnt++
		b = b.Next
	}
	if cnt < k {
		return a
	}
	newHead := reverseBetween(head, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

//翻转区间[start,end)的链表
func reverseBetween(start, end *collections.ListNode) *collections.ListNode {
	if start == nil || start == end || start.Next == end {
		return start
	}
	var pre *collections.ListNode
	cur := start
	for cur != end {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

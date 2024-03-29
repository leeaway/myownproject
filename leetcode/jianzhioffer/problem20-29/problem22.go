package problem20_29

import (
	"example.com/m/v2/tools/collections"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/21 20:00
 */

// 输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//
// 例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
//
//
//
// 示例：
//
//
//给定一个链表: 1->2->3->4->5, 和 k = 2.
//
//返回链表 4->5.
//
// Related Topics 链表 双指针 👍 412 👎 0

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//快慢指针，让fast先走k步，slow再出发，等fast到底后slow在的位置就是倒数第k个
func getKthFromEnd(head *collections.ListNode, k int) *collections.ListNode {
	slow, fast := head, head
	cnt := 0
	for fast != nil {
		fast = fast.Next
		if cnt < k {
			cnt++
			continue
		}
		slow = slow.Next
	}
	return slow
}

package LinkedList

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/28 20:40
 */

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
// 你应当 保留 两个分区中每个节点的初始相对位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [2,1], x = 2
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200
//
//
// Related Topics 链表 双指针 👍 650 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *collections.ListNode, x int) *collections.ListNode {
	less, more := &collections.ListNode{}, &collections.ListNode{}
	dummyLess, dummyMore := less, more
	for head != nil {
		if head.Val < x {
			dummyLess.Next = &collections.ListNode{Val: head.Val}
			dummyLess = dummyLess.Next
		} else {
			dummyMore.Next = &collections.ListNode{Val: head.Val}
			dummyMore = dummyMore.Next
		}
		head = head.Next
	}
	dummyLess.Next = more.Next
	return less.Next
}

//leetcode submit region end(Prohibit modification and deletion)

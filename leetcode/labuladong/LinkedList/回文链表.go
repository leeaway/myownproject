package LinkedList

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/30 18:52
 */

//给定一个链表的 头节点 head ，请判断其是否为回文链表。
//
// 如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
//
//
//
// 示例 1：
//
//
//
//
//输入: head = [1,2,3,3,2,1]
//输出: true
//
// 示例 2：
//
//
//
//
//输入: head = [1,2]
//输出: false
//
//
//
//
// 提示：
//
//
// 链表 L 的长度范围为 [1, 10⁵]
// 0 <= node.val <= 9
//
//
//
//
// 进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
//
//
//
//
// 注意：本题与主站 234 题相同：https://leetcode-cn.com/problems/palindrome-linked-list/
//
// Related Topics 栈 递归 链表 双指针 👍 88 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
  方法1: 正常方法比较简单，就是遍历一遍链表，然后存到数组里，利用首尾指针判断即可，但是空间复杂度O(n)
  方法2: 考虑空间O(1)的做法，先快慢指针找到链表中点，然后将后半段链表反转，对比反转后的即可；
*/
func isPalindrome(head *collections.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	midNode := getMidNode(head)
	reverseMid := reverseNode(midNode)
	return checkIsPalindrome(head, reverseMid)
}

func checkIsPalindrome(before, after *collections.ListNode) bool {
	for after != nil {
		if before.Val != after.Val {
			return false
		}
		before = before.Next
		after = after.Next
	}
	return true
}

func getMidNode(head *collections.ListNode) *collections.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	res := slow.Next
	//断开原链表
	slow.Next = nil
	return res
}

func reverseNode(node *collections.ListNode) *collections.ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	var pre *collections.ListNode
	cur := node
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)

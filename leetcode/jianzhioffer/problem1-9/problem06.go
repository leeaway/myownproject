package problem1_9

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/10/12 10:03
 */

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
//
//
//
// 示例 1：
//
// 输入：head = [1,3,2]
//输出：[2,3,1]
//
//
//
// 限制：
//
// 0 <= 链表长度 <= 10000
//
// Related Topics 栈 递归 链表 双指针 👍 342 👎 0

/*
方法一：迭代
	每次append到头部即可
*/
func reversePrint1(head *collections.ListNode) []int {
	var res []int
	for head != nil {
		res = append([]int{head.Val}, res...)
		head = head.Next
	}
	return res
}

/*
 方法一的递归写法
*/

func reversePrint1_recurse(head *collections.ListNode) []int {
	var res []int
	if head == nil {
		return res
	}
	res = append([]int{head.Val}, res...)
	return append(reversePrint1_recurse(head.Next), res...)
}

/*
	递归
*/

func reversePrint2(head *collections.ListNode) []int {
	return reversePrint2helper(head, []int{})
}

func reversePrint2helper(head *collections.ListNode, res []int) []int {
	if head == nil {
		return res
	}
	res = reversePrint2helper(head.Next, res)
	res = append(res, head.Val)
	return res
}

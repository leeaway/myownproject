package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 15:15
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Show() []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

package HotOneHundred

//两数之和，非O(n^2)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		val, ok := numMap[target-num]
		if ok {
			return []int{val, i}
		}
		numMap[num] = i
	}
	return nil
}

// 两数相加
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	cur := &ListNode{}
	res := cur
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return res.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//进位记录
	upper := 0
	dummy := &ListNode{Val: 0}
	res := dummy
	for l1 != nil && l2 != nil {
		curVal := (l1.Val + l2.Val + upper) % 10
		if l1.Val+l2.Val+upper > 9 {
			upper = 1
		} else {
			upper = 0
		}
		dummy.Next = &ListNode{Val: curVal}
		dummy = dummy.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		curVal := (l1.Val + upper) % 10
		if l1.Val+upper > 9 {
			upper = 1
		} else {
			upper = 0
		}
		dummy.Next = &ListNode{Val: curVal}
		dummy = dummy.Next
		l1 = l1.Next
	}

	for l2 != nil {
		curVal := (l2.Val + upper) % 10
		if l2.Val+upper > 9 {
			upper = 1
		} else {
			upper = 0
		}
		dummy.Next = &ListNode{Val: curVal}
		dummy = dummy.Next
		l2 = l2.Next
	}

	if upper > 0 {
		dummy.Next = &ListNode{
			Val: upper,
		}
	}

	return res.Next
}

//反转链表
func reverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

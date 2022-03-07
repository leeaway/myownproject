package HotOneHundred

import "example.com/m/v2/tools/mathutil"

/*
	001，两数之和，非O(n^2)
	思路：前缀和
*/

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

// 002，两数相加
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

/*
  003，无重复字符的最长子串
	思路：滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	l, r, res := 0, 1, 1
	digitMap := make(map[rune]bool)
	digitMap[rune(s[l])] = true
	for r < len(s) {
		for r < len(s) && !digitMap[rune(s[r])] {
			digitMap[rune(s[r])] = true
			r++
		}
		res = mathutil.MaxInt(res, r-l)
		if r == len(s) {
			return res
		}
		for s[l] != s[r] {
			delete(digitMap, rune(s[l]))
			l++
		}
		for l < r && s[l] == s[r] {
			delete(digitMap, rune(s[l]))
			l++
		}
		digitMap[rune(s[r])] = true
		r++
	}
	return res
}

/*
	hot004,寻找两个正序数组的中位数,要求O(log(m+n))
	思路: 二分法，找到第（m+n）/2个数
	主要思路：要找到第 k (k>1) 小的元素，那么就取 pivot1 = nums1[k/2-1] 和 pivot2 = nums2[k/2-1] 进行比较 * 这里的 "/" 表示整除
	1.nums1 中小于等于 pivot1 的元素有     nums1[0 .. k/2-2] 共计 k/2-1 个
	2.nums2 中小于等于 pivot2 的元素有 nums2[0 .. k/2-2] 共计 k/2-1 个
	3.取 pivot = min(pivot1, pivot2)，两个数组中小于等于 pivot 的元素共计不会超过 (k/2-1) + (k/2-1) <= k-2 个
	4.这样 pivot 本身最大也只能是第 k-1 小的元素
	5.如果 pivot = pivot1，那么 nums1[0 .. k/2-1] 都不可能是第 k 小的元素。把这些元素全部 "删除"，剩下的作为新的 nums1 数组
	6.如果 pivot = pivot2，那么 nums2[0 .. k/2-1] 都不可能是第 k 小的元素。把这些元素全部 "删除"，剩下的作为新的 nums2 数组
	7.由于我们 "删除" 了一些元素（这些元素都比第 k 小的元素要小），因此需要修改 k 的值，减去删除的数的个数

*/

//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	m,n := len(nums1),len(nums2)
//	l1,r1,l2,r2 := 0,m-1,0,n-1
//
//}

/*
	hot005,最长回文子串
	思路一：中心扩散法
*/

func longestPalindrome(s string) string {
	res := s[:1]
	for i := 0; i < len(s)-1; i++ {
		cur1 := check(s, i, i)
		cur2 := check(s, i, i+1)
		if len(cur1) > len(res) {
			res = cur1
		}
		if len(cur2) > len(res) {
			res = cur2
		}
	}
	return res
}

func check(str string, center1, center2 int) string {
	for center1 >= 0 && center2 < len(str) && str[center1] == str[center2] {
		center1--
		center2++
	}
	return str[center1+1 : center2]
}

/*
	hot006,正则表达式匹配
	动态规划：dp[i][j]:s的前i个字符与p的前j个字符是否匹配
	状态转移：
*/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	return true
}

func matchHelper(a, b rune) bool {
	if a == b || b == '.' {
		return true
	}
	return false
}

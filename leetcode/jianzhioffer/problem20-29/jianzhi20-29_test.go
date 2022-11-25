package problem20_29

import (
	"example.com/m/v2/tools/collections"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 15:17
 */

func Test21(t *testing.T) {
	fmt.Println(exchange([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(exchange([]int{1, 4, 2, 6, 5, 9, 8, 7}))
}

func Test_22(t *testing.T) {
	Convey("链表中倒数第k个节点", t, func() {
		Convey("22 test1", func() {
			node4 := &collections.ListNode{
				Val:  4,
				Next: nil,
			}
			node3 := &collections.ListNode{
				Val:  3,
				Next: node4,
			}

			node2 := &collections.ListNode{
				Val:  2,
				Next: node3,
			}

			head := &collections.ListNode{
				Val:  1,
				Next: node2,
			}
			So(getKthFromEnd(head, 2), ShouldEqual, node3)
		})
	})
}

func Test_24(t *testing.T) {
	Convey("反转链表", t, func() {
		Convey("24 test1", func() {
			node4 := &collections.ListNode{
				Val:  4,
				Next: nil,
			}
			node3 := &collections.ListNode{
				Val:  3,
				Next: node4,
			}

			node2 := &collections.ListNode{
				Val:  2,
				Next: node3,
			}

			head := &collections.ListNode{
				Val:  1,
				Next: node2,
			}
			newHead := reverseList(head)
			So(newHead.Val, ShouldEqual, 4)
			So(newHead.Next.Val, ShouldEqual, 3)
			So(newHead.Next.Next.Val, ShouldEqual, 2)
			So(newHead.Next.Next.Next.Val, ShouldEqual, 1)
			So(newHead.Next.Next.Next.Next, ShouldEqual, nil)
		})
	})
}

func Test_25(t *testing.T) {
	Convey("合并两个排序的链表", t, func() {
		Convey("25 test1", func() {
			node4 := &collections.ListNode{
				Val:  4,
				Next: nil,
			}
			node3 := &collections.ListNode{
				Val:  3,
				Next: node4,
			}

			node2 := &collections.ListNode{
				Val:  2,
				Next: node3,
			}

			head1 := &collections.ListNode{
				Val:  1,
				Next: node2,
			}

			node7 := &collections.ListNode{
				Val:  7,
				Next: nil,
			}
			node6 := &collections.ListNode{
				Val:  5,
				Next: node7,
			}

			node5 := &collections.ListNode{
				Val:  3,
				Next: node6,
			}

			head2 := &collections.ListNode{
				Val:  1,
				Next: node5,
			}
			mergeList := mergeTwoLists(head1, head2)
			mergeList2 := mergeTwoLists2(head1, head2)
			So(mergeList.Show(), ShouldResemble, []int{1, 1, 2, 3, 3, 4, 5, 7})
			So(mergeList2.Show(), ShouldResemble, []int{1, 1, 2, 3, 3, 4, 5, 7})
		})
	})
}

func Test_29(t *testing.T) {
	Convey("spiralOrder", t, func() {
		Convey("n*n螺旋矩阵 test1", func() {
			matrix := [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			}
			So(spiralOrder(matrix), ShouldResemble, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
		})

		Convey("m*n螺旋矩阵", func() {
			matrix := [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
			}
			So(spiralOrder(matrix), ShouldResemble, []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6})
		})

		Convey("空矩阵", func() {
			matrix := [][]int{}
			So(spiralOrder(matrix), ShouldResemble, []int{})
		})
	})
}

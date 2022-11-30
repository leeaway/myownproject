package LinkedList

import (
	"container/heap"
	"example.com/m/v2/tools/collections"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/28 20:46
 */

func Test_86(t *testing.T) {
	Convey("分隔链表", t, func() {
		Convey("分隔链表", func() {
			node6 := &collections.ListNode{
				Val:  2,
				Next: nil,
			}
			node5 := &collections.ListNode{
				Val:  5,
				Next: node6,
			}
			node4 := &collections.ListNode{
				Val:  2,
				Next: node5,
			}
			node3 := &collections.ListNode{
				Val:  3,
				Next: node4,
			}

			node2 := &collections.ListNode{
				Val:  4,
				Next: node3,
			}

			head1 := &collections.ListNode{
				Val:  1,
				Next: node2,
			}

			So(partition(head1, 2).Show(), ShouldResemble, []int{1, 4, 3, 2, 5, 2})
			So(partition(head1, 3).Show(), ShouldResemble, []int{1, 3, 2, 2, 4, 5})
		})
	})
}

func Test_MergeKListNode(t *testing.T) {
	Convey("MergeKListNode", t, func() {
		Convey("MergeKListNode test1", func() {
			node1 := collections.NewListNodeByArray([]int{1, 4, 7, 10, 13})
			node2 := collections.NewListNodeByArray([]int{2, 5, 8, 11})
			node3 := collections.NewListNodeByArray([]int{3, 6, 9, 12, 15})
			node4 := collections.NewListNodeByArray([]int{1, 1, 2, 2, 3, 3})
			Convey("minListNodeHeap test", func() {
				minListNodeHeap := &ListNodeHeap{node3, node1, node2}
				heap.Init(minListNodeHeap)
				So(heap.Pop(minListNodeHeap).(*collections.ListNode).Val, ShouldEqual, 1)
				So(heap.Pop(minListNodeHeap).(*collections.ListNode).Val, ShouldEqual, 2)
				So(heap.Pop(minListNodeHeap).(*collections.ListNode).Val, ShouldEqual, 3)
			})
			So(mergeKLists([]*collections.ListNode{node1, node2, node3}).Show(), ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 15})
			So(mergeKLists([]*collections.ListNode{node2, node3, node4}).Show(), ShouldResemble, []int{1, 1, 2, 2, 2, 3, 3, 3, 5, 6, 8, 9, 11, 12, 15})
		})
	})
}

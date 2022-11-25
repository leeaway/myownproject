package problem11_19

import (
	"example.com/m/v2/tools/collections"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 15:21
 */

func Test10(t *testing.T) {
	Convey("斐波那契", t, func() {
		So(fib_dp(100), ShouldEqual, 687995182)
		So(fib_recursive(40), ShouldEqual, fib_dp(40))
	})
}

func Test11(t *testing.T) {
	Convey("旋转数组的最小值", t, func() {
		So(minArray([]int{3, 4, 5, 1, 2}), ShouldEqual, 1)
		So(minArray([]int{2, 2, 2, 2, 2, 2}), ShouldEqual, 2)
		So(minArray([]int{3, 4, 5, 6, 3, 3, 3}), ShouldEqual, 3)
		So(minArray([]int{3, 4, 5, 6, 2, 3, 3}), ShouldEqual, 2)

		So(minArray2([]int{3, 4, 5, 1, 2}), ShouldEqual, 1)
		So(minArray2([]int{2, 2, 2, 2, 2, 2}), ShouldEqual, 2)
		So(minArray2([]int{3, 4, 5, 6, 3, 3, 3}), ShouldEqual, 3)
		So(minArray2([]int{3, 4, 5, 6, 2, 3, 3}), ShouldEqual, 2)
	})
}

func Test12(t *testing.T) {
	Convey("矩阵中的路径", t, func() {
		board1 := [][]byte{
			{'A', 'B', 'C', 'D'},
			{'D', 'E', 'F', 'G'},
			{'C', 'E', 'C', 'K'},
			{'D', 'F', 'B', 'A'},
		}
		board2 := [][]byte{
			{'C', 'A', 'A'},
			{'A', 'A', 'A'},
			{'B', 'C', 'D'},
		}
		So(exist(board1, "CFCB"), ShouldEqual, true)
		So(exist(board1, "BCFCBFD"), ShouldEqual, true)
		So(exist(board1, "CDCB"), ShouldEqual, false)
		So(exist(board2, "AAB"), ShouldEqual, true)
	})
}

func Test14_1(t *testing.T) {
	Convey("剪绳子I", t, func() {
		So(cuttingRope(3), ShouldEqual, 2)
		So(cuttingRope(4), ShouldEqual, 4)
		So(cuttingRope(10), ShouldEqual, 36)
	})
}

func Test15(t *testing.T) {
	Convey("二进制中1的个数", t, func() {
		So(hammingWeight2(11), ShouldEqual, 3)
		So(hammingWeight2(128), ShouldEqual, 1)
		So(hammingWeight1(11), ShouldEqual, 3)
		So(hammingWeight1(128), ShouldEqual, 1)
	})
}

func Test16(t *testing.T) {
	Convey("数值的整数次方", t, func() {
		So(myPow(2.00000, 10), ShouldEqual, 1024.00000)
		So(myPow(2.00000, -2), ShouldEqual, 0.25000)
	})
}

func Test17(t *testing.T) {
	Convey("打印从1到最大的n位数", t, func() {
		So(printNumbers(1), ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})
}

func Test18(t *testing.T) {
	Convey("删除链表节点", t, func() {
		node3 := &collections.ListNode{
			Val:  3,
			Next: nil,
		}
		node2 := &collections.ListNode{
			Val:  2,
			Next: node3,
		}
		node1 := &collections.ListNode{
			Val:  1,
			Next: node2,
		}
		head := &collections.ListNode{
			Val:  0,
			Next: node1,
		}
		So(deleteNode(head, 1).Show(), ShouldResemble, []int{0, 2, 3})
	})
}

package jianzhioffer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/10/11 16:38
 */

func Test03(t *testing.T) {
	Convey("找出重复数字", t, func() {
		So(findRepeatNumber1([]int{1, 3, 4, 5, 6, 2, 3}), ShouldEqual, 3)
		So(findRepeatNumber2([]int{1, 3, 4, 5, 6, 2, 3}), ShouldEqual, 3)
	})
}

func Test04(t *testing.T) {
	Convey("查找二维数组", t, func() {
		matrix := [][]int{
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
			{26, 27, 28, 29, 30},
		}
		So(findNumberIn2DArray1(matrix, 11), ShouldEqual, true)
		So(findNumberIn2DArray1(matrix, 23), ShouldEqual, true)
		So(findNumberIn2DArray1(matrix, 9), ShouldEqual, false)
		So(findNumberIn2DArray2(matrix, 11), ShouldEqual, true)
		So(findNumberIn2DArray2(matrix, 23), ShouldEqual, true)
		So(findNumberIn2DArray2(matrix, 9), ShouldEqual, false)
	})
}

func Test05(t *testing.T) {
	Convey("替换空格", t, func() {
		So(replaceSpace("hello world"), ShouldEqual, "hello%20world")
	})
}

func Test06(t *testing.T) {
	node3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	head := &ListNode{
		Val:  0,
		Next: node1,
	}
	Convey("从尾到头打印联表", t, func() {
		So(reversePrint1(head), ShouldResemble, []int{3, 2, 1, 0})
		So(reversePrint1_recurse(head), ShouldResemble, []int{3, 2, 1, 0})
		So(reversePrint2(head), ShouldResemble, []int{3, 2, 1, 0})
	})
}

func Test07(t *testing.T) {
	Convey("重建二叉树", t, func() {
		Convey("左右子树都有", func() {
			preorder := []int{3, 9, 20, 15, 7}
			inorder := []int{9, 3, 15, 20, 7}
			tree := buildTree(preorder, inorder)
			So(preOrder(tree, []int{}), ShouldResemble, preorder)
			So(inOrder(tree, []int{}), ShouldResemble, inorder)
		})

		Convey("没有左子树", func() {
			preorder2 := []int{3, 20, 15, 7}
			inorder2 := []int{3, 15, 20, 7}
			tree2 := buildTree(preorder2, inorder2)
			So(preOrder(tree2, []int{}), ShouldResemble, preorder2)
			So(inOrder(tree2, []int{}), ShouldResemble, inorder2)
		})

		Convey("没有右子树", func() {
			preorder2 := []int{3, 10, 9, 11}
			inorder2 := []int{9, 10, 11, 3}
			tree2 := buildTree(preorder2, inorder2)
			So(preOrder(tree2, []int{}), ShouldResemble, preorder2)
			So(inOrder(tree2, []int{}), ShouldResemble, inorder2)
		})
	})
}

func Test09(t *testing.T) {
	Convey("两个栈实现队列", t, func() {
		cQue := Constructor()
		cQue.AppendTail(1)
		cQue.AppendTail(2)
		So(cQue.DeleteHead(), ShouldEqual, 1)
		cQue.AppendTail(3)
		cQue.AppendTail(4)
		So(cQue.DeleteHead(), ShouldEqual, 2)
		So(cQue.DeleteHead(), ShouldEqual, 3)
		cQue.AppendTail(5)
		So(cQue.DeleteHead(), ShouldEqual, 4)
		So(cQue.DeleteHead(), ShouldEqual, 5)
	})
}

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

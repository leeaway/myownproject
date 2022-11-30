package problem30_39

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 15:10
 */

func Test_30(t *testing.T) {
	Convey("包含min的栈", t, func() {
		Convey("30 test1", func() {
			minStack := MinStackConstructor()
			minStack.Push(-2)
			minStack.Push(0)
			minStack.Push(-3)
			So(minStack.Min(), ShouldEqual, -3)
			minStack.Pop()
			So(minStack.Min(), ShouldEqual, -2)
			minStack.Pop()
			So(minStack.Min(), ShouldEqual, -2)
			minStack.Pop()
			minStack.Push(4)
			So(minStack.Min(), ShouldEqual, 4)

		})
	})
}

func Test_31(t *testing.T) {
	Convey("栈的压入弹出顺序校验", t, func() {
		Convey("31 test1", func() {
			So(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}), ShouldEqual, true)
			So(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}), ShouldEqual, true)
			So(validateStackSequences([]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 3, 2, 1, 6}), ShouldEqual, true)
		})
	})
}

func Test_33(t *testing.T) {
	Convey("是否是二叉搜索树的后序遍历", t, func() {
		Convey("33 test1", func() {
			So(verifyPostorder([]int{1, 3, 2, 6, 5}), ShouldEqual, true)
			So(verifyPostorder([]int{1, 3, 2, 5, 6}), ShouldEqual, true)
			So(verifyPostorder([]int{1, 6, 3, 2, 5}), ShouldEqual, false)
			So(verifyPostorder([]int{7, 4, 6, 5}), ShouldEqual, false)

		})
	})
}

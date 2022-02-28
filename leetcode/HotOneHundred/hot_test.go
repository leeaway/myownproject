package HotOneHundred

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHot001(t *testing.T) {
	Convey("TestTwoSum", t, func() {
		So(twoSum([]int{2, 7, 11, 15}, 9), ShouldResemble, []int{0, 1})
		So(twoSum([]int{3, 3, 4, 10}, 6), ShouldResemble, []int{0, 1})
		So(twoSum([]int{3, 2, 4, 11, 15}, 6), ShouldResemble, []int{1, 2})
	})
}

func TestHot002(t *testing.T) {
	Convey("TestAddTwoNum", t, func() {
		So(addTwoNumbers(NewListNode([]int{2, 4, 3}), NewListNode([]int{5, 6, 4})), ShouldResemble, NewListNode([]int{7, 0, 8}))
		So(addTwoNumbers(NewListNode([]int{9, 9, 9, 9, 9, 9, 9}), NewListNode([]int{9, 9, 9, 9})), ShouldResemble, NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}))
	})
}

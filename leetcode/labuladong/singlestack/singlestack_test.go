package singlestack

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/29 18:14
 */

func Test_(t *testing.T) {
	Convey("nextGreatNum", t, func() {
		Convey("nextGreatNum test1", func() {
			So(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}), ShouldResemble, []int{-1, 3, -1})
			So(nextGreaterElement([]int{4, 1, 2}, []int{1, 2, 3, 4}), ShouldResemble, []int{-1, 2, 3})
		})
	})

	Convey("nextLessNum", t, func() {
		Convey("nextLessNum test1", func() {
			So(nextLessElement([]int{1, 3, 4, 2}), ShouldResemble, []int{-1, 2, 2, -1})
			So(nextLessElement([]int{1, 2, 3, 4}), ShouldResemble, []int{-1, -1, -1, -1})
		})
	})
}

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
	Convey("", t, func() {
		Convey(" test1", func() {
			So(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}), ShouldResemble, []int{-1, 3, -1})
		})
	})
}

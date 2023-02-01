package binarysearch

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/1 17:52
 */

func Test_binarySearchTarget(t *testing.T) {
	Convey("binarySearchTarget", t, func() {
		Convey("binarySearchTarget test1", func() {
			So(binarySearchTarget([]int{1, 2, 3, 5, 6, 8}, 2), ShouldEqual, 1)
		})
		Convey("left bound", func() {
			So(binarySearchTargetLeftBound([]int{1, 2, 2, 2, 2, 3, 3, 8}, 2), ShouldEqual, 1)
		})
		Convey("right bound", func() {
			So(binarySearchTargetRightBound([]int{1, 2, 2, 2, 2, 3, 3, 8}, 2), ShouldEqual, 4)
		})
	})
}

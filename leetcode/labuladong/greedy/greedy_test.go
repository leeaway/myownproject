package greedy

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 11:52
 */

func Test_MinOperation(t *testing.T) {
	Convey("MinOperation", t, func() {
		Convey("MinOperation test1", func() {
			So(minOperations([]int{1, 1, 2}, []int{4}), ShouldEqual, 0)
			So(minOperations([]int{1, 3, 6, 1, 2, 5, 1}, []int{2}), ShouldEqual, -1)
			So(minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}), ShouldEqual, 3)
			So(minOperations([]int{5, 2, 1, 5, 2, 2, 2, 2, 4, 3, 3, 5}, []int{1, 4, 5, 5, 6, 3, 1, 3, 3}), ShouldEqual, 1)
		})
	})
}

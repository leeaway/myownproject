package singlequeue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/9 15:06
 */

/*
	单调队列用于求取区间的最大或最小值
	注意：单调队列在算法应用中大多是基于双端队列(Deque)实现的，因为经常会涉及其他操作，而不是单单维护一个单调队列。。
*/

func Test_getMaxInSlidingWindow(t *testing.T) {
	Convey("getMaxInSlidingWindow", t, func() {
		Convey("getMaxInSlidingWindow test1", func() {
			So(getMaxInSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), ShouldResemble, []int{3, 3, 5, 5, 6, 7})

			So(getMaxInSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3), ShouldResemble, []int{3, 3, 2, 5})
		})
	})
}

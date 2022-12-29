package problem110_119

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/29 11:52
 */

func Test_longestConsecutive(t *testing.T) {
	Convey("", t, func() {
		Convey(" test1", func() {
			So(longestConsecutive([]int{1, 2, 100, 200, 4, 3, 90, 1, 2}), ShouldEqual, 4)
			So(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}), ShouldEqual, 9)
		})
	})
}

func Test_findRedundantConnection(t *testing.T) {
	Convey("", t, func() {
		Convey(" test1", func() {
			So(findRedundantConnection([][]int{
				{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5},
			}), ShouldResemble, []int{1, 4})
		})
	})
}

func Test_findOrder(t *testing.T) {
	Convey("findOrder", t, func() {
		Convey("findOrder test1", func() {
			fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
		})
	})
}

func Test_sequenceReconstruction(t *testing.T) {
	Convey("sequenceReconstruction", t, func() {
		Convey("sequenceReconstruction test1", func() {
			So(sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}}), ShouldEqual, false)
			So(sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}), ShouldEqual, true)
		})
	})
}

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
	Convey("test03", t, func() {
		So(findRepeatNumber1([]int{1, 3, 4, 5, 6, 2, 3}), ShouldEqual, 3)
		So(findRepeatNumber2([]int{1, 3, 4, 5, 6, 2, 3}), ShouldEqual, 3)
	})
}

func Test04(t *testing.T) {
	Convey("test04", t, func() {
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
	Convey("test05", t, func() {
		So(replaceSpace("hello world"), ShouldEqual, "hello%20world")
	})
}

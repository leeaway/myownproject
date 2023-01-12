package problem60_69

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/3 18:48
 */

func Test_(t *testing.T) {
	Convey("", t, func() {
		Convey(" test1", func() {
			So(constructArr([]int{2, 3, 4, 5, 6}), ShouldResemble, []int{360, 240, 180, 144, 120})
			So(constructArr([]int{1, 2, 3, 4, 5}), ShouldResemble, []int{120, 60, 40, 30, 24})
		})
	})
}

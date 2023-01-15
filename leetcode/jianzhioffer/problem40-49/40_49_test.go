package problem40_49

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/13 23:09
 */

func Test_getLeastNumbers(t *testing.T) {
	Convey("getLeastNumbers", t, func() {
		Convey("getLeastNumbers test1", func() {
			So(getLeastNumbers([]int{1, 3, 5, 6, 2, 4}, 2), ShouldResemble, []int{2, 1})
			So(getLeastNumbers([]int{1, 3, 5, 6, 2, 4}, 3), ShouldResemble, []int{3, 2, 1})
		})
	})
}

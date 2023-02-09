package slidewindow

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/1 18:06
 */

func Test_minWindow(t *testing.T) {
	Convey("minWindow", t, func() {
		Convey("minWindow test1", func() {
			So(minWindow("ABKEGCHJBANC", "ABC"), ShouldEqual, "BANC")
			So(minWindow2("ABKEGCHJBANC", "ABC"), ShouldEqual, "BANC")
		})
	})
}

func Test_checkInclusion(t *testing.T) {
	Convey("checkInclusion", t, func() {
		Convey("checkInclusion test1", func() {
			So(checkInclusion("ab", "ebodbacf"), ShouldEqual, true)
		})
	})
}

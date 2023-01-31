package futu

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/30 17:13
 */

func Test_checkRepeatedStr(t *testing.T) {
	Convey("checkRepeatedStr", t, func() {
		Convey("checkRepeatedStr test1", func() {
			So(checkRepeatedStr("abcabc"), ShouldEqual, true)
			So(checkRepeatedStr("ababa"), ShouldEqual, false)
			So(checkRepeatedStr("aaaaab"), ShouldEqual, false)
			So(checkRepeatedStr("aaaaa"), ShouldEqual, true)
		})
	})
}

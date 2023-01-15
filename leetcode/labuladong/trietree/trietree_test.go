package trietree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/15 14:18
 */

func Test_(t *testing.T) {
	Convey("", t, func() {
		Convey(" test1", func() {
			magic := Constructor()
			magic.BuildDict([]string{"abc", "abd", "bcf"})
			So(magic.Search("abc"), ShouldEqual, true)
			So(magic.Search("abcd"), ShouldEqual, false)
		})
	})
}

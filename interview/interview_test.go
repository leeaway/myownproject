package interview

import (
	"example.com/m/v2/interview/futu"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/16 14:57
 */

func Test_Solve(t *testing.T) {
	Convey("Solve", t, func() {
		Convey("Solve test1", func() {
			res := futu.Solve([]int{0, 1})
			fmt.Println(res)

		})
	})
}

func Test_ConcurrentPrint123(t *testing.T) {
	Convey("ConcurrentPrint123", t, func() {
		Convey("ConcurrentPrint123 test1", func() {
			futu.ConcurrentPrint123()
		})
	})
}

func Test_getRemainNum(t *testing.T) {
	Convey("getRemainNum", t, func() {
		Convey("getRemainNum test1", func() {
			So(getRemainNum("123421", 2), ShouldEqual, 1221)
			So(getRemainNum("10300", 1), ShouldEqual, 300)
		})
	})
}

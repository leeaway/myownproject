package mathutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/11/22 11:18
 */

func Test_GCD(t *testing.T) {
	Convey("GCD", t, func() {
		Convey("GCD test1", func() {
			So(Gcd(12, 28), ShouldEqual, 4)
			So(Gcd(7, 4), ShouldEqual, 1)
			So(Gcd(96, 64), ShouldEqual, 32)
		})
	})
}

func Test_Lcm(t *testing.T) {
	Convey("Lcm", t, func() {
		Convey("Lcm test1", func() {
			So(Lcm(12, 6), ShouldEqual, 12)
			So(Lcm(6, 4), ShouldEqual, 12)
			So(Lcm(4, 7), ShouldEqual, 28)
			So(Lcm(120, 80), ShouldEqual, 240)
		})
	})
}

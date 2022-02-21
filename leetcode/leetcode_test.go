package leetcode

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLeetcode688(t *testing.T) {
	Convey("TestLeetcode688", t, func() {
		So(KnightProbability2(3, 2, 0, 0), ShouldEqual, 0.0625)
		So(KnightProbability2(1, 0, 0, 0), ShouldEqual, float64(1))
		So(KnightProbability2(3, 2, 1, 1), ShouldEqual, float64(0))
	})
}

func TestLeetcode838(t *testing.T) {
	Convey("TestLeeCode838", t, func() {
		So(PushDominoes("RR.L"), ShouldEqual, "RR.L")
		So(PushDominoes(".L.R...LR..L.."), ShouldEqual, "LL.RR.LLRRLL..")
	})
}

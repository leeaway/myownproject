package Huawei

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/9 21:47
 */

func Test_Probelm1(t *testing.T) {
	Convey("Probelm1", t, func() {
		Convey("Probelm1 test1", func() {
			fmt.Println(Base16ToBase2("05"))
			//So(solveHelper("E980A5"), ShouldEqual, 36901)
			//So(solveHelper("C0C0"), ShouldEqual, -1)
		})
	})
}

func Test_Probelm2(t *testing.T) {
	Convey("Probelm2", t, func() {
		Convey("Probelm2 test1", func() {
			So(solveProblem2([]int{1, 1, 1}, 4), ShouldEqual, 0)
		})
	})
}

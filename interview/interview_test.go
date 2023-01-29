package interview

import (
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
			res := solve([]int{0, 1})
			fmt.Println(res)

		})
	})
}

func Test_ConcurrentPrint123(t *testing.T) {
	Convey("ConcurrentPrint123", t, func() {
		Convey("ConcurrentPrint123 test1", func() {
			ConcurrentPrint123()
		})
	})
}

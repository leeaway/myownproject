package concurent

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/9 15:54
 */

func Test_StartDemo(t *testing.T) {
	Convey("StartDemo", t, func() {
		Convey("StartDemo test1", func() {
			StartDemo()
		})
	})
}

func Test_Print0To100(t *testing.T) {
	Convey("Print0To100", t, func() {
		Convey("Print0To100 test1", func() {
			ConcurrentPrint0To100(4)
		})
	})
}

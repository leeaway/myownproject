package sortMethod

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSort(t *testing.T) {
	Convey("Test Bubble Sort", t, func() {
		So(bubbleSort([]int{1, 2, 4, 9, 3, 5, 8, 7, 6}), ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})
}

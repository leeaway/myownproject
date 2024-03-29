package sortMethod

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSort(t *testing.T) {
	//测试冒泡排序
	Convey("Test Bubble Sort", t, func() {
		So(bubbleSort([]int{1, 2, 4, 9, 3, 5, 8, 7, 6}), ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	//测试插入排序
	Convey("Test Insert Sort", t, func() {
		So(insertSort([]int{1, 2, 4, 9, 3, 5, 8, 7, 6}), ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	})

	//测试快速排序
	Convey("Test Quick Sort", t, func() {
		nums := []int{1, 2, 4, 2, 9, 3, 5, 2, 8, 7, 6}
		quickSort(nums)
		So(nums, ShouldResemble, []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9})
	})
}

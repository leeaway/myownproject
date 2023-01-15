package collections

import (
	"container/heap"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/13 22:43
 */

func Test_IntHeap(t *testing.T) {
	Convey("IntHeap", t, func() {
		Convey("IntHeap 大根堆", func() {
			maxHeap := NewMaxIntHeap()
			heap.Init(maxHeap)
			heap.Push(maxHeap, 3)
			heap.Push(maxHeap, 2)
			So(maxHeap.Peek().(int), ShouldEqual, 3)

			heap.Push(maxHeap, 5)
			So(maxHeap.Peek().(int), ShouldEqual, 5)
			So(heap.Pop(maxHeap).(int), ShouldEqual, 5)
			So(maxHeap.Peek().(int), ShouldEqual, 3)
			So(heap.Pop(maxHeap).(int), ShouldEqual, 3)
			So(maxHeap.Peek().(int), ShouldEqual, 2)
			heap.Push(maxHeap, 1)
			So(heap.Pop(maxHeap), ShouldEqual, 2)
			So(heap.Pop(maxHeap), ShouldEqual, 1)
		})

		Convey("IntHeap 小根堆", func() {
			minHeap := NewMinIntHeap()
			heap.Init(minHeap)
			heap.Push(minHeap, 3)
			heap.Push(minHeap, 2)
			So(minHeap.Peek(), ShouldEqual, 2)
			So(heap.Pop(minHeap), ShouldEqual, 2)
			heap.Push(minHeap, 1)
			So(minHeap.Peek(), ShouldEqual, 1)
			So(heap.Pop(minHeap), ShouldEqual, 1)
			So(heap.Pop(minHeap), ShouldEqual, 3)
		})
	})
}

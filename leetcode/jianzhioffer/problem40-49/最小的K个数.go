package problem40_49

import (
	"container/heap"
	"example.com/m/v2/tools/collections"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/13 23:09
 */

//输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
//
//
// 示例 1：
//
// 输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
//
//
// 示例 2：
//
// 输入：arr = [0,1,2,1], k = 1
//输出：[0]
//
//
//
// 限制：
//
//
// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 507 👎 0

/*
	用一个最大堆即可
*/
//leetcode submit region begin(Prohibit modification and deletion)
func getLeastNumbers(arr []int, k int) []int {
	maxHeap := collections.NewMaxIntHeap()
	heap.Init(maxHeap)
	for _, a := range arr {
		if maxHeap.Size() == k {
			if maxHeap.Peek().(int) > a {
				heap.Pop(maxHeap)
				heap.Push(maxHeap, a)
			}
			continue
		}
		heap.Push(maxHeap, a)
	}
	var res []int
	for !maxHeap.IsEmpty() {
		res = append(res, heap.Pop(maxHeap).(int))
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

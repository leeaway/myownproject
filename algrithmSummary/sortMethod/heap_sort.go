package sortMethod

/**
 * @author 2416144794@qq.com
 * @date 2023/2/7 14:08
 */
/*
	算法描述：首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，将 其中最大的值变为父节点，递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整O(nlgn).

	算法步骤，以小顶堆为例
	1. 创建堆，根据数组索引映射到完全二叉树即可
	2. 调整堆,从无序部分最后一个非叶子节点开始，每次比较是否要跟两个子节点交换，注意交换后可能还要跟子节点对比，这里递归实现，每次将最大值调整到队首；
	3. 交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)，尾节点为有序
*/

//arr 存储堆的数组
//n 数组长度
//i 待维护元素的下标
func heapify(arr []int, n int, i int) {
	max := i
	//左孩子节点
	lson := 2*i + 1
	//右孩子节点
	rson := 2*i + 2
	if lson < n && arr[lson] > arr[max] {
		max = lson
	}
	if rson < n && arr[rson] > arr[max] {
		max = rson
	}
	//把父节点换下去并向下调整
	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		heapify(arr, n, max)
	}
}

func HeapSort(arr []int, n int) {
	//建堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	//排序
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

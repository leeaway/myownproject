package sortMethod

/*
	快速排序，时间复杂度O(N^2),不稳定
	主要思想：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序
	比如：每次以第一个元素为base，将数组分为两部分，左边都不大于base，右边都大于base。
	具体做法：
		1，创建两个指针分别指向数组的最左端以及最右端
		2，在数组中任意取出一个元素作为基准（以第一个元素为例）
		3，左指针开始向右移动，遇到比基准大的停止
		4，右指针开始向左移动，遇到比基准小的元素停止，交换左右指针所指向的元素
		5，重复3，4，直到左指针超过右指针，此时，比基准小的值就都会放在基准的左边，比基准大的值会出现在基准的右边
		6，然后分别对基准的左右两边重复以上的操作，直到数组完全排序
*/
func quickSort(nums []int, start, end int) []int {
	if len(nums) < 2 {
		return nums
	}
	base := nums[0]
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l] <= base {
			l++
		}
		for r > l && nums[r] >= base {
			r--
		}
		nums[l], nums[r] = swap(nums[l], nums[r])
	}
	nums[0] = nums[l]
	nums[l] = base
	leftPart := quickSort(nums[0:l])
	rightPart := quickSort(nums[l+1:])
	return append(append(leftPart, base), rightPart...)
}

func swap(a, b int) (int, int) {
	return b, a
}

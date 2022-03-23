package sortMethod

/*
	插入排序，时间复杂度O(N^2),稳定排序
	类似摸牌后的整理
*/
func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		j := i
		//这里类似插牌，这里可以优化成二分法，因为前面的数一定是有序的
		for j > 0 && nums[j] < nums[j-1] {
			tmp := nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = tmp
			j--
		}
	}
	return nums
}

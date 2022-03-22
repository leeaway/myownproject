package sortMethod

/*
	冒泡排序，时间复杂度O(N^2),稳定排序
*/
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//交换
			if nums[j] < nums[i] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	return nums
}

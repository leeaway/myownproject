package dailyExercise

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		for r >= 0 && (nums[r]%2) == 1 {
			r--
		}
		for l < n && (nums[l]%2) == 0 {
			l++
		}
		if l < r {
			tmp := nums[l]
			nums[l] = nums[r]
			nums[r] = tmp
		}
	}
	return nums
}

package dailyExercise

func numSubarrayProductLessThanK(nums []int, k int) int {
	n, multi := len(nums), 1
	res, l := 0, 0
	for j := 0; j < n; j++ {
		multi *= nums[j]
		for l <= j && multi >= k {
			multi /= nums[l]
			l++
		}
		res += j - l + 1
	}
	return res
}

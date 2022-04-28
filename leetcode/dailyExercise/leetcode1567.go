package dailyExercise

func getMaxLen(nums []int) int {
	n := len(nums)
	pos, neg := make([]int, n), make([]int, n) //表示以nums[i]结尾的乘积为正数和负数的最长子数组长度
	pos[0], neg[0] = 0, 0
	if nums[0] > 0 {
		pos[0] = 1
	}
	if nums[0] < 0 {
		neg[0] = 1
	}
	res := pos[0]
	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			if neg[i-1] > 0 {
				pos[i] = neg[i-1] + 1
			}
			neg[i] = pos[i-1] + 1
		} else if nums[i] > 0 {
			pos[i] = pos[i-1] + 1
			if neg[i-1] > 0 {
				neg[i] = neg[i-1] + 1
			}
		} else {
			pos[i] = 0
			neg[i] = 0
		}
		res = max(res, pos[i])
	}
	return res
}

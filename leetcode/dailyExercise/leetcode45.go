package dailyExercise

func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		tmp := 10001
		for j := i + 1; j <= min(i+nums[i], n-1); j++ {
			tmp = min(tmp, dp[j])
		}
		dp[i] = tmp + 1
	}
	return dp[0]
}

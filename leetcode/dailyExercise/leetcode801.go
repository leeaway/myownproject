package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/10/10 11:38
 */

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = n + 1
		}
	}
	dp[0][0], dp[0][1] = 0, 1
	for i := 1; i < n; i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			//前一个也不换
			dp[i][0] = min(dp[i][0], dp[i-1][0])
			//或者前一个换了，这个也换
			dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
			//或者前一个换了，但是这个不换，此时需满足
			if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
				dp[i][0] = min(dp[i][0], dp[i-1][1])
			}
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			//此时前一个不换，当前换
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
			//或者前一个换了，这个不换
			dp[i][0] = min(dp[i][0], dp[i-1][1])
			//或者前一个换了，这个也换，那就要满足
			if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
				dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
			}
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 14:37
 */

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// Related Topics 数组 二分查找 动态规划 👍 2898 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
动态规划，
	定义：dp[i]表示以nums[i]结尾的最长子序列
	状态转移：dp[i] = max(dp[k]+1,dp[i]),其中k满足，0<=i<k && nums[k]<nums[i]
	Base Case: dp[0] = 1,dp[k]默认值为1,因为本身就是一个子序列
	最后遍历dp，返回最大值即可
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = mathutil.MaxInt(dp[i], dp[j]+1)
			}
		}
	}
	res := 1
	for _, d := range dp {
		if d > res {
			res = d
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

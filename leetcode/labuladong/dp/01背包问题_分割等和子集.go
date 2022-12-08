package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/8 10:31
 */

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
// 示例 2：
//
//
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 100
//
//
// Related Topics 数组 动态规划 👍 1578 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	本质上就是求出target = sum/2,然后从nums中选出n个数，使得和等于target
	转化为背包问题：给一个可装载重量为 target 的背包和 N 个物品，每个物品的重量为 nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满？
	定义：dp[n][target+1],其中dp[i][j]表示nums[:i]是否存在一种装法将j容量的背包装满；
	状态转移：dp[i][j] = dp[i][j] || dp[i-1][j] || dp[i-1][j-nums[i]],其中nums[i]<j
			其中：dp[i-1][j]:表示不拿当前的nums[i],则dp[i][j]依赖dp[i-1][j]
			     dp[i-1][j-nums[i]]: 表示拿当前的nums[i]，则依赖nums[:i-1]中是否能恰好装满j-nums[i]的背包
	Base Case：
			1. dp[0][nums[0]] = true,其他dp[0][k] = false
			2. dp[k][0] = false,因为所有num都大于0

*/
func canPartition(nums []int) bool {
	sum := mathutil.Sum(nums)
	n := len(nums)
	if sum%2 > 0 || n < 2 {
		return false
	}
	target := sum / 2
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		if nums[0] <= target {
			dp[0][nums[0]] = true
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i][j] || dp[i-1][j]
			if j > nums[i] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)

package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/8 11:38
 */

//给你一个整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
//
//
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3
//
//
// 示例 2：
//
//
//输入：nums = [1], target = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000
//
//
// Related Topics 数组 动态规划 回溯 👍 1451 👎 0

/*
	我们假设nums中正数之和为x，负数之和为y，则有：
		x+y = sum
		x-y = target
	解方程可得：x = (sum+target)/2,所以若sum + target 不是偶数或者小于0，直接返回0；
	反之，问题则变成了与Leetcode416 分割等和子集类似了，不过这里是计数，我们统一用背包问题解决；
    值得注意的是，nums中0元素队结果无影响，所以每含有一个0，结果就要乘于2，然后真正计算只保留大于0得数进行计算

	定义：dp[n][t+1]，其中dp[i][j]表示nums[:i+1]中可以搭配出值为j的表达式个数,t为(sum+target)/2+1
	状态转移：不选当前： dp[i][j] = dp[i-1][j],
			选当前： dp[i][j] += dp[i-1][j-nums[i]]
	Base Case:
			1. dp[0][0] = 1,因为可以不选nums[0],target为0；
			2. dp[0][nums[0]] = 1,其他dp[0][k]=0，注意关注nums[0]与t的大小关系，注意是否越界
*/

//leetcode submit region begin(Prohibit modification and deletion)
func findTargetSumWays(nums []int, target int) int {
	var newNums []int
	multi := 1
	for _, num := range nums {
		if num > 0 {
			newNums = append(newNums, num)
			continue
		}
		multi *= 2
	}
	return findTargetSumWaysHelper(newNums, target) * multi
}
func findTargetSumWaysHelper(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	sum := mathutil.Sum(nums)
	newS := sum + target
	if newS < 0 || newS%2 > 0 {
		return 0
	}
	target = newS / 2
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
		dp[0][0] = 1
		if nums[0] <= target {
			dp[0][nums[0]] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] += dp[i-1][j]
			if j >= nums[i] {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}

//leetcode submit region end(Prohibit modification and deletion)

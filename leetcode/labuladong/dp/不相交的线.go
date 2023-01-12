package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2023/1/11 22:56
 */

//在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
//
//
// nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
//
//
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
// 以这种方法绘制线条，并返回可以绘制的最大连线数。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,4,2], nums2 = [1,2,4]
//输出：2
//解释：可以画出两条不交叉的线，如上图所示。
//但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相
//交。
//
//
//
// 示例 2：
//
//
//
//输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
//输出：3
//
//
//
// 示例 3：
//
//
//
//输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
//输出：2
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 500
// 1 <= nums1[i], nums2[j] <= 2000
//
//
//
//
// Related Topics 数组 动态规划 👍 404 👎 0

/*
	动态规划做法：
	定义：dp[i][j],以nums1[i-1],nums2[j-1]结尾的最大连线数
	转移：
		1. nums1[i] == nums2[j],则dp[i][j] = max(dp[i-1][j-1]+1,dp[i-1][j],dp[i][j-1]),ij连线或不连线
		2. nums1[i] != nums2[j],则dp[i][j] = max(dp[i-1][j],dp[i][j-1])
	Base Case:
		dp[0][j] = dp[i][0] = 0
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = mathutil.MaxInt(dp[i-1][j], dp[i][j-1])
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = mathutil.MaxInt(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}

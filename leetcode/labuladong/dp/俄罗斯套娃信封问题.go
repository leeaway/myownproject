package dp

import (
	"example.com/m/v2/tools/mathutil"
	"sort"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 15:08
 */

//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 注意：不允许旋转信封。
//
// 示例 1：
//
//
//输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出：3
//解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//
// 示例 2：
//
//
//输入：envelopes = [[1,1],[1,1],[1,1]]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= envelopes.length <= 10⁵
// envelopes[i].length == 2
// 1 <= wi, hi <= 10⁵
//
//
// Related Topics 数组 二分查找 动态规划 排序 👍 839 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
动态规划：
	定义：dp[i]表示以第i个信封为最外层的套娃层数；
	状态转移：dp[i] = max(dp[i],dp[k]+1),其中0<=k<i && envelopes[k]可以放在envelopes[i]中；
	Base Case: dp[i] = 1
	边界条件：这里需要保证当i>j时，envelope[i] 一定不能放在 envelope[j]中，否者上述的状态转移矩阵无效，因此需要对envelopes先按宽度，宽度相同时按高度升序排序均可；
*/
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	res := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if canCover(envelopes[j], envelopes[i]) {
				dp[i] = mathutil.MaxInt(dp[i], dp[j]+1)
				res = mathutil.MaxInt(dp[i], res)
			}
		}
	}
	return res
}

//检查1是否能放在2中
func canCover(envelope1, envelope2 []int) bool {
	return envelope2[0] > envelope1[0] && envelope2[1] > envelope1[1]
}

//leetcode submit region end(Prohibit modification and deletion)

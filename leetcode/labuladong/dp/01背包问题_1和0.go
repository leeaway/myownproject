package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/8 14:59
 */

//给你一个二进制字符串数组 strs 和两个整数 m 和 n 。
//
//
// 请你找出并返回 strs 的最大子集的长度，该子集中 最多 有 m 个 0 和 n 个 1 。
//
//
// 如果 x 的所有元素也是 y 的元素，集合 x 是集合 y 的 子集 。
//
//
//
// 示例 1：
//
//
//输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
//输出：4
//解释：最多有 5 个 0 和 3 个 1 的最大子集是 {"10","0001","1","0"} ，因此答案是 4 。
//其他满足题意但较小的子集包括 {"0001","1"} 和 {"10","1","0"} 。{"111001"} 不满足题意，因为它含 4 个 1 ，大于
//n 的值 3 。
//
//
// 示例 2：
//
//
//输入：strs = ["10", "0", "1"], m = 1, n = 1
//输出：2
//解释：最大的子集是 {"0", "1"} ，所以答案是 2 。
//
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 600
// 1 <= strs[i].length <= 100
// strs[i] 仅由 '0' 和 '1' 组成
// 1 <= m, n <= 100
//
//
// Related Topics 数组 字符串 动态规划 👍 829 👎 0

/*
	这题也是典型的背包问题，限定0和1的数量上限，求出最多能拿多少个元素
	定义：dp[l][m+1][n+1],dp[i][j][k]表示strs[:i+1]中限制j个0，k个1上限的情况下最多能拿多少个元素
	转移矩阵： 不拿当前：dp[i][j][k] = dp[i-1][j][k]
			 拿当前：  dp[i][j][k] = dp[i-1][j-c0][k-c1]+1,c0,c1表示strs[i]的0，1个数
	Base Case: dp[0][c0~m][c1~n] = 1,即选用第一个元素，则必然要求上限要满足；
*/
//leetcode submit region begin(Prohibit modification and deletion)
func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	indexToCountMap := buildMap(strs)
	dp := make([][][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	for j := indexToCountMap[0][0]; j <= m; j++ {
		for k := indexToCountMap[0][1]; k <= n; k++ {
			dp[0][j][k] = 1
		}
	}

	for i := 1; i < l; i++ {
		c0, c1 := indexToCountMap[i][0], indexToCountMap[i][1]
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = mathutil.MaxInt(dp[i][j][k], dp[i-1][j][k])
				if j >= c0 && k >= c1 {
					dp[i][j][k] = mathutil.MaxInt(dp[i][j][k], dp[i-1][j-c0][k-c1]+1)
				}
			}
		}
	}

	return dp[l-1][m][n]
}

func buildMap(strs []string) map[int][]int {
	res := make(map[int][]int)
	for i, str := range strs {
		res[i] = count(str)
	}
	return res
}

func count(str string) []int {
	res := make([]int, 2)
	for _, s := range str {
		if s == '0' {
			res[0]++
			continue
		}
		res[1]++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

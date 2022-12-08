package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 20:37
 */

//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//
//
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
//
//
// 示例 1：
//
//
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace" ，它的长度为 3 。
//
//
// 示例 2：
//
//
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc" ，它的长度为 3 。
//
//
// 示例 3：
//
//
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= text1.length, text2.length <= 1000
// text1 和 text2 仅由小写英文字符组成。
//
//
// Related Topics 字符串 动态规划 👍 1178 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	动态规划：
	定义：dp[i][j] 表示text1[:i],text2[:j]的最长公共子序列，i,j 范围为[0,m],[0,n]
	状态转移：若text1[i] == text2[j],则：dp[i][j] = dp[i-1][j-1]+1
			反之：dp[i][j] = max(dp[i-1][j],dp[i-1][j-1],dp[i][j-1])
	Base Case: dp[0][k]:text1无字符，公共子串肯定为0，dp[k][0]同理
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				continue
			}
			dp[i][j] = mathutil.MaxInt(dp[i-1][j], mathutil.MaxInt(dp[i-1][j-1], dp[i][j-1]))
		}
	}
	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)

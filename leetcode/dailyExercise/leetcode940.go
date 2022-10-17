package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/10/14 11:17
 */

//给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。
//
// 字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。
//
//
// 例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。
//
//
//
//
// 示例 1：
//
//
//输入：s = "abc"
//输出：7
//解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
//
//
// 示例 2：
//
//
//输入：s = "aba"
//输出：6
//解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
//
//
// 示例 3：
//
//
//输入：s = "aaa"
//输出：3
//解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
//
//
// Related Topics 字符串 动态规划 👍 210 👎 0

/*
方法：动态规划
定义dp[i][j]:表示截止到i位置以j结尾的不同序列个数，其中0<=i<n,0<=j<26(代表26个字母)
则： 若s[i] != j,则i位置不选，即dp[i][j] = dp[i-1][j]
	若s[i] == j,则i位置可选，即dp[i][j] = sum(dp[i-1][k])+1,其中0<=k<26
		+1是因为s[i]可以作为一个单独的序列
*/

func distinctSubseqII(s string) int {
	if len(s) == 0 {
		return 0
	}
	n, mod := len(s), 1000000007
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 26)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < 26; j++ {
			if int(s[i-1]-'a') != j {
				dp[i][j] = dp[i-1][j]
			} else {
				cur := 1
				for k := 0; k < 26; k++ {
					cur = (cur + dp[i-1][k]) % mod
				}
				dp[i][j] = cur
			}
		}
	}
	ans := 0
	for i := 0; i < 26; i++ {
		ans = (ans + dp[n][i]) % mod
	}
	return ans
}

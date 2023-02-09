package dp

/**
 * @author 2416144794@qq.com
 * @date 2023/2/2 18:23
 */

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//
// 示例 2：
//
//
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//
// 示例 3：
//
//
//输入：s = ""
//输出：0
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 3 * 10⁴
// s[i] 为 '(' 或 ')'
//
//
// Related Topics 栈 字符串 动态规划 👍 2136 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	定义：dp[i]表示以s[i]结尾的最长有效括号的长度，显然若s[i] = '(',结果为0
	转移：
		1. s[i]=')'&&s[i-1]='(',则有dp[i] = dp[i-2]+2
		2. s[i]=')'&&s[i-1]=')':
			若s[i-dp[i-2]-1]='(',则有dp[i] = dp[i-1]+dp[i-dp[i-2]-2]+2
		3. 其他情况dp[i] = 0
	Base：初始化为0即可
*/
func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n)
	res := 0
	for i := 2; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else {
			if s[i-dp[i-2]-1] == '(' {
				dp[i] = dp[i-1] + dp[i-dp[i-2]-2] + 2
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

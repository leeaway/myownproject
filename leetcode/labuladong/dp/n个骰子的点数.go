package dp

/**
 * @author 2416144794@qq.com
 * @date 2023/1/29 11:12
 */

//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
//
//
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
//
//
//
// 示例 1:
//
// 输入: 1
//输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
//
//
// 示例 2:
//
// 输入: 2
//输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0
//.05556,0.02778]
//
//
//
// 限制：
//
// 1 <= n <= 11
//
// Related Topics 数学 动态规划 概率与统计 👍 510 👎 0

/*
	定义：dp[i][j]: i个骰子点数之和为j的情况个数
	转移：dp[i][j] = sum(dp[i-1][j-k]), 1<= k <=6
	边界：
		dp[1][1~6] = 1，其他都为0
*/

//leetcode submit region begin(Prohibit modification and deletion)
func dicesProbability(n int) []float64 {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 6*n+1)
		for j := 1; j <= 6 && i == 1; j++ {
			dp[1][j] = 1
		}
	}

	for i := 2; i <= n; i++ {
		for j := i; j <= 6*i; j++ {
			for k := 1; k <= 6 && j > k; k++ {
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}

	sum := 0
	for j := n; j <= 6*n; j++ {
		sum += dp[n][j]
	}

	var res []float64
	for j := n; j <= 6*n; j++ {
		res = append(res, float64(dp[n][j])/float64(sum))
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

package dp

import (
	"example.com/m/v2/tools/mathutil"
	"fmt"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 17:39
 */

//给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作：
//
//
// 插入一个字符
// 删除一个字符
// 替换一个字符
//
//
//
//
// 示例 1：
//
//
//输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//
//
// 示例 2：
//
//
//输入：word1 = "intention", word2 = "execution"
//输出：5
//解释：
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
//
//
//
//
// 提示：
//
//
// 0 <= word1.length, word2.length <= 500
// word1 和 word2 由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 2705 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
动态规划：类似这种两个字符串的，一般都是双指针，i,j分别对应word1,word2
	定义：dp[i][j]标识word1[:i],word2[:j]匹配上的最小编辑距离,m,n := len(word1),len(word2),0<=i<=m,0<=j<=n
	状态转移：编辑方法有三种，加上什么都不操作，其实有四种，有：
		1.若word1[i] == word2[j],不编辑，i++,j++
		2.word1插入一个和word2[j]相同的字符，j++,操作+1
		  word1删除当前的字符，i++,操作+1
		  word1替换成word2[j],i++,j++,操作+1
	Base Case:  dp[0][k]:表示word1没有字符，word2有k个字符，显然需要k次操作；
			    dp[k][0]:同理
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = i
		} else {
			for j := 1; j <= n; j++ {
				dp[0][j] = j
			}
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			add := dp[i][j-1] + 1
			del := dp[i-1][j] + 1
			rep := dp[i-1][j-1] + 1
			dp[i][j] = mathutil.MinInt(add, mathutil.MinInt(del, rep))
		}
	}
	return dp[m][n]
}

/*
方法二：其实本质上是可以递归的
*/

func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return m + n
	}
	if word1[0] == word2[0] {
		return minDistance2(word1[1:], word2[1:])
	}
	add := minDistance2(word1[1:], word2) + 1
	del := minDistance2(word1, word2[1:]) + 1
	rep := minDistance2(word1[1:], word2[1:]) + 1
	return mathutil.MinInt(add, mathutil.MinInt(del, rep))
}

/*
	显然上面的递归解法有很多的重复子问题，可能会超时，我们尝试加上备忘录
	本质上记录两个指针，以两个指针作key即可
*/

func minDistance3(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return m + n
	}
	memo := make(map[string]int)
	return minDistanceHelper(word1, word2, 0, 0, memo)
}

func buildKeyStr(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}

func minDistanceHelper(word1 string, word2 string, p1, p2 int, memo map[string]int) int {
	val, ok := memo[buildKeyStr(p1, p2)]
	if ok {
		return val
	}

	m, n := len(word1), len(word2)
	if m == p1 || n == p2 {
		return m + n - (p1 + p2)
	}

	if word1[p1] == word2[p2] {
		memo[buildKeyStr(p1, p2)] = minDistanceHelper(word1, word2, p1+1, p2+1, memo)
	} else {
		add := minDistanceHelper(word1, word2, p1+1, p2, memo) + 1
		del := minDistanceHelper(word1, word2, p1, p2+1, memo) + 1
		rep := minDistanceHelper(word1, word2, p1+1, p2+1, memo) + 1
		memo[buildKeyStr(p1, p2)] = mathutil.MinInt(add, mathutil.MinInt(del, rep))
	}
	return memo[buildKeyStr(p1, p2)]
}

//leetcode submit region end(Prohibit modification and deletion)

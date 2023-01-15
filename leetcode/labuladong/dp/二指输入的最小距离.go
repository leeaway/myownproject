package dp

import (
	"example.com/m/v2/tools/mathutil"
	"math"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/1/13 15:54
 */

//
//
// 二指输入法定制键盘在 X-Y 平面上的布局如上图所示，其中每个大写英文字母都位于某个坐标处。
//
//
// 例如字母 A 位于坐标 (0,0)，字母 B 位于坐标 (0,1)，字母 P 位于坐标 (2,3) 且字母 Z 位于坐标 (4,1)。
//
//
// 给你一个待输入字符串 word，请你计算并返回在仅使用两根手指的情况下，键入该字符串需要的最小移动总距离。
//
// 坐标 (x1,y1) 和 (x2,y2) 之间的 距离 是 |x1 - x2| + |y1 - y2|。
//
// 注意，两根手指的起始位置是零代价的，不计入移动总距离。你的两根手指的起始位置也不必从首字母或者前两个字母开始。
//
//
//
// 示例 1：
//
//
//输入：word = "CAKE"
//输出：3
//解释：
//使用两根手指输入 "CAKE" 的最佳方案之一是：
//手指 1 在字母 'C' 上 -> 移动距离 = 0
//手指 1 在字母 'A' 上 -> 移动距离 = 从字母 'C' 到字母 'A' 的距离 = 2
//手指 2 在字母 'K' 上 -> 移动距离 = 0
//手指 2 在字母 'E' 上 -> 移动距离 = 从字母 'K' 到字母 'E' 的距离  = 1
//总距离 = 3
//
//
// 示例 2：
//
//
//输入：word = "HAPPY"
//输出：6
//解释：
//使用两根手指输入 "HAPPY" 的最佳方案之一是：
//手指 1 在字母 'H' 上 -> 移动距离 = 0
//手指 1 在字母 'A' 上 -> 移动距离 = 从字母 'H' 到字母 'A' 的距离 = 2
//手指 2 在字母 'P' 上 -> 移动距离 = 0
//手指 2 在字母 'P' 上 -> 移动距离 = 从字母 'P' 到字母 'P' 的距离 = 0
//手指 1 在字母 'Y' 上 -> 移动距离 = 从字母 'A' 到字母 'Y' 的距离 = 4
//总距离 = 6
//
//
//
//
// 提示：
//
//
// 2 <= word.length <= 300
// 每个 word[i] 都是一个大写英文字母。
//
//
// Related Topics 字符串 动态规划 👍 81 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	定义: dp[i][j][k]表示当前是第i轮，第一根手指在j处，第二根手指在k处,其中 0<=i<len(word),0<=j,k<26
	转移：当前是第i轮，对应字母为word[i],记为w，则有：
		1. 当前轮左手移到w位置，dp[i][w][R] = min(dp[i][w][R],dp[i-1][L][R]+cal(L,w))
		2. 当前轮右手移到w位置，dp[i][L][w] = min(dp[i][L][w],dp[i-1][L][R]+cal(R,w))

	Base Case:
		我们知道不是所有状态都是合法的，所以我们先默认给所有值付一个MaxInt值，这样对最优解就没有影响
		当然因为初始位置不算距离：
			左手在word[0]时，dp[0][word[0]-'a'][k] = 0
			右手在word[0]时，dp[0][j][word[0]-'a'] = 0
*/
func minimumDistance(word string) int {
	n := len(word)
	dp := make([][][]int, n+1)
	//初始化
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 26)
		for j := 0; j < 26; j++ {
			dp[i][j] = make([]int, 26)
			for k := 0; k < 26; k++ {
				if i > 0 {
					dp[i][j][k] = math.MaxInt
				}
			}
		}
	}

	res := math.MaxInt
	for i := 1; i <= n; i++ {
		w := int(word[i-1] - 'A')
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				if dp[i-1][j][k] == math.MaxInt {
					continue
				}
				dp[i][w][k] = mathutil.MinInt(dp[i][w][k], dp[i-1][j][k]+calDistance(j, w))
				if i == n {
					res = mathutil.MinInt(res, dp[i][w][k])
					res = mathutil.MinInt(res, dp[i][j][w])
				}
			}
		}
	}
	return res
}

func calDistance(a, b int) int {
	tmp := int(math.Abs(float64(a) - float64(b)))
	return tmp/6 + tmp%6
}

//leetcode submit region end(Prohibit modification and deletion)

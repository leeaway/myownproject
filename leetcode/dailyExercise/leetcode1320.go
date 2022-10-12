package dailyExercise

import "math"

/**
 * @author 2416144794@qq.com
 * @date 2022/10/12 14:37
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
// Related Topics 字符串 动态规划 👍 77 👎 0

/*
方法：动态规划，本题着手点就是左右手的位置
	dp[i][j]:表示最终左手在i位置，右手在j位置,i,j一定不相等，因为两根手指不能在同一个位置
	由于对称性，直接设i>j，则有状态转移：
	1.若i-j > 1，则dp[i][j] = dp[i-1][j]+dis(i-1,i),因为此时右手在j<i-1位置，则i-1位置一定是左手点击的；
	2.若i-j = 1，则左手来源可以是[0,i-1),即dp[i][j] = min(dp[k][j]+dis(k,i)),其中k的范围为[0,i-1)；
	初始条件：dp[i][0] = sum(dis(k,k+1)),其中k的范围为[1,i-1]
*/

func minimumDistance(word string) int {
	n := len(word)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 2; i < n; i++ {
		dp[i][0] = dp[i-1][0] + calDisBetweenTwoRunes(i-1, i, word)
	}

	for j := 1; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			if i-j > 1 {
				dp[i][j] = dp[i-1][j] + calDisBetweenTwoRunes(i-1, i, word)
			} else {
				dp[i][j] = math.MaxInt32
				for k := 0; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i-1][k]+calDisBetweenTwoRunes(k, i, word))
				}
			}
		}
	}

	res := math.MaxInt32
	for j := 0; j < n-1; j++ {
		res = min(res, dp[n-1][j])
	}
	return res
}

func calDisBetweenTwoRunes(i, j int, word string) int {
	a, b := word[i], word[j]
	diff := int(math.Abs(float64(a) - float64(b)))
	return diff/6 + diff%6
}

package dp

import "example.com/m/v2/tools/mathutil"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/12 18:05
 */

//
//
// 一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿
//过地下城并通过对抗恶魔来拯救公主。
//
// 骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
//
// 有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），要么
//包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
//
// 为了尽快到达公主，骑士决定每次只向右或向下移动一步。
//
//
//
// 编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。
//
// 例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。
//
//
//
// -2 (K) -3  3
// -5    -10  1
// 10    30  -5 (P)
//
// 说明:
//
//
// 骑士的健康点数没有上限。
// 任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
//
//
// Related Topics 数组 动态规划 矩阵 👍 682 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

/*
方法：我们从公主（P）出发去找王子，更好计算营救公主的最低血量,如下图graph:
	-2 (K) -3  3
	-5    -10  1
	10    30  -5 (P)
注意点：
1.王子血量必须 >= 1,=0代表死亡；
2.王子只能向右或向下走，对应公主即向左向上；

因此，我们用dp代表最少到达此地至少需要的血量，初始为（其中#代表整数最大值）：
	#   #   #       #   #   #        #   #   2          7   5   2
	#   #   #  =》  #   #   5   =》  #   11   5   =》   6   11   5
	#   #   6      #   1   6        1   1   6          1   1   6

其中需要注意的是，dp[1][1]的计算，我们到达graph[1][2]至少需要5滴血（因为这里会加1滴血，所以到P时有6滴血，刚好满足），所以：
1.往右走，此时需要5-(-10) = 15滴血；
2.往下走，此时需要1-(-10) = 11滴血
所以dp[1][1] = 11;

另外：dp[2][1]的计算，graph[2][1]会加30滴血，所以按6-30 = -24算，王子已经go die了，所以此时应该是按最低血量来，即dp[2][1] = -1

动态规划：
定义： dp[m][n],dp[i][j]代表要救出Princess，走到该点所需的最低血量；
状态转移:
	1. 0<=i<m-1 && 0<=j<n-1: dp[i][j] = min(max(dp[i+1][j] - graph[i][j],1),max(dp[i][j+1]-graph[i][j],1))
	2. 0<=i<m-1 && j==n-1:     dp[i][j] = max(dp[i+1][j] - graph[i][j],1)
	3. i==m-1 && 0<=j<n-1:   dp[i][j] = max(dp[i][j+1]-graph[i][j],1)
base case:
	1. graph[m-1][n-1] >= 0, dp[m-1][n-1] = 1;
	2. graph[m-1][n-1] < 0,  dp[m-1][n-1] = 1-graph[m-1][n-1]
*/
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[m-1][n-1] = 1
	if dungeon[m-1][n-1] < 0 {
		dp[m-1][n-1] = 1 - dungeon[m-1][n-1]
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j < n-1 {
				dp[i][j] = mathutil.MaxInt(dp[i][j+1]-dungeon[i][j], 1)
				continue
			}
			if j == n-1 && i < m-1 {
				dp[i][j] = mathutil.MaxInt(dp[i+1][j]-dungeon[i][j], 1)
				continue
			}
			if i < m-1 && j < n-1 {
				dp[i][j] = mathutil.MinInt(mathutil.MaxInt(dp[i+1][j]-dungeon[i][j], 1), mathutil.MaxInt(dp[i][j+1]-dungeon[i][j], 1))
			}
		}
	}
	return dp[0][0]
}

//leetcode submit region end(Prohibit modification and deletion)

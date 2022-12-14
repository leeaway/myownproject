package dp

import (
	"example.com/m/v2/tools/mathutil"
	"math"
)

/**
 * @author 2416144794@qq.com
 * @date 2022/12/12 21:09
 */

//有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城
//市 fromi 开始，以价格 pricei 抵达 toi。
//
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便
//宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
//
//
// 示例 1：
//
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 1
//输出: 200
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
//
// 示例 2：
//
//
//输入:
//n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//src = 0, dst = 2, k = 0
//输出: 500
//解释:
//城市航班图如下
//
//
//从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 10⁴
// 航班没有重复，且不存在自环
// 0 <= src, dst, k < n
// src != dst
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 动态规划 最短路 堆（优先队列） 👍 543 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	动态规划：
	定义：dp[n][n][k],dp[i][j][l]表示从城市i到城市j的最便宜机票价格,剩余l次中转机会，默认值为Integer.MAX_VALUE
	状态转移：dp[i][j][l] = min(dp[z][j][l-1]+price(z,j)),其中z为所有可以到达j的城市,price(z,j)代表从z到j的航班的价格,用二维数组来记录；
	Base Case:
		1. dp[i][j][0],不能中转，则 dp[i][j][0] = price(i,j)
*/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	canReachNodes := make(map[int][]int)
	prices := make([][]int, n)
	for i := 0; i < n; i++ {
		prices[i] = make([]int, n)
	}
	for _, flight := range flights {
		from, to, price := flight[0], flight[1], flight[2]
		canReachNodes[to] = append(canReachNodes[to], from)
		prices[from][to] = price
	}
	return 0
}

//代表从from到to的最低价格
func findCheapestPriceHelper(canReachNodes map[int][]int, prices [][]int, from, to, k int) int {
	if from == to {
		return 0
	}
	cur := math.MaxInt
	for _, node := range canReachNodes[to] {
		cur = mathutil.MinInt(findCheapestPriceHelper(canReachNodes, prices, from, node, k-1), cur)
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)

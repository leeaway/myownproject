package backtrack

import "strings"

/**
 * @author 2416144794@qq.com
 * @date 2023/1/11 23:24
 */

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 1606 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	基本思路：
	逐行放皇后，采用回溯的方法，穷举所有可能，保存可行的方案
*/
var res [][]string

func solveNQueens(n int) [][]string {
	graph := make([][]string, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]string, n)
		for j := 0; j < n; j++ {
			graph[i][j] = "."
		}
	}
	solveNQueensHelper(graph, 0)
	return res
}

func solveNQueensHelper(graph [][]string, row int) {
	m, n := len(graph), len(graph[0])
	if row == m {
		var tmp []string
		for _, item := range graph {
			tmp = append(tmp, strings.Join(item, ""))
		}
		res = append(res, tmp)
		return
	}

	for j := 0; j < n; j++ {
		//在第j列放Q
		if checkCanPut(graph, row, j) {
			graph[row][j] = "Q"
			solveNQueensHelper(graph, row+1)
			graph[row][j] = "."
		}
	}
}

//在graph的[row,col]位置放置一个Q，是否可行,时间复杂度O(n*n)
func checkCanPut(graph [][]string, row, col int) bool {
	for i := 0; i < row; i++ {
		for j := 0; j < len(graph[0]); j++ {
			if graph[i][j] == "Q" {
				if i == row || j == col || (row-i) == (col-j) || row+col-i-j == 0 {
					return false
				}
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

package problem11_19

/**
 * @author 2416144794@qq.com
 * @date 2022/10/12 21:50
 */

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。
//
//
//
//
//
// 示例 1：
//
//
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word =
//"ABCCED"
//输出：true
//
//
// 示例 2：
//
//
//输入：board = [["a","b"],["c","d"]], word = "abcd"
//输出：false
//
//
//
//
// 提示：
//
//
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成
//
//
// 注意：本题与主站 79 题相同：https://leetcode-cn.com/problems/word-search/
//
// Related Topics 数组 回溯 矩阵 👍 686 👎 0

// 这是典型的回溯题，我们只需要在board中寻找word[0],然后以此为起点去找路径是否存在即可
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}
			if findPath(board, used, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func findPath(board [][]byte, used [][]bool, x, y, wordIndex int, word string) bool {
	m, n := len(board), len(board[0])
	if board[x][y] != word[wordIndex] {
		return false
	}
	if wordIndex == len(word)-1 {
		return true
	}
	used[x][y] = true
	dx, dy := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	res := false
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx < 0 || nx >= m || ny < 0 || ny >= n || used[nx][ny] {
			continue
		}
		used[nx][ny] = true
		res = res || findPath(board, used, nx, ny, wordIndex+1, word)
		used[nx][ny] = false
	}
	used[x][y] = false
	return res
}

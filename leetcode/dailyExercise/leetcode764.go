package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/11/9 10:29
 */

//在一个 n x n 的矩阵 grid 中，除了在数组 mines 中给出的元素为 0，其他每个元素都为 1。mines[i] = [xi, yi]表示
//grid[xi][yi] == 0
//
// 返回 grid 中包含 1 的最大的 轴对齐 加号标志的阶数 。如果未找到加号标志，则返回 0 。
//
// 一个 k 阶由 1 组成的 “轴对称”加号标志 具有中心网格 grid[r][c] == 1 ，以及4个从中心向上、向下、向左、向右延伸，长度为 k-1，
//由 1 组成的臂。注意，只有加号标志的所有网格要求为 1 ，别的网格可能为 0 也可能为 1 。
//
//
//
// 示例 1：
//
//
//
//
//输入: n = 5, mines = [[4, 2]]
//输出: 2
//解释: 在上面的网格中，最大加号标志的阶只能是2。一个标志已在图中标出。
//
//
// 示例 2：
//
//
//
//
//输入: n = 1, mines = [[0, 0]]
//输出: 0
//解释: 没有加号标志，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 500
// 1 <= mines.length <= 5000
// 0 <= xi, yi < n
// 每一对 (xi, yi) 都 不重复
//
//
// Related Topics 数组 动态规划 👍 135 👎 0

/*
	暴力解法：
*/
func orderOfLargestPlusSign(n int, mines [][]int) int {
	if n == 1 {
		return 0
	}
	graph := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			graph[i][j] = true
		}
	}

	for _, index := range mines {
		graph[index[0]][index[1]] = false
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !graph[i][j] {
				continue
			}
			res = max(res, calRadius(graph, i, j))
		}
	}
	return res
}

func calRadius(graph [][]bool, x, y int) int {
	n := len(graph)
	up, down, l, r := x, x, y, y
	ur, dr, lr, rr := 0, 0, 0, 0
	for up >= 0 && graph[up][y] {
		up--
		ur++
	}
	for down < n && graph[down][y] {
		down++
		dr++
	}

	for l >= 0 && graph[x][l] {
		l--
		lr++
	}

	for r < n && graph[x][r] {
		r++
		rr++
	}
	return min(min(ur, dr), min(lr, rr))
}
package leetcode

func findBall(grid [][]int) []int {
	n := len(grid[0])
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, solve(grid, i))
	}
	return res
}

func solve(grid [][]int, cur int) int {
	col := cur
	m := len(grid)
	row := 0
	for row < m && !isStuck(grid, row, col) {
		if grid[row][col] == 1 {
			row++
			col++
			continue
		}
		row++
		col--
	}
	if row != m {
		return -1
	}
	return col
}

func isStuck(grid [][]int, row, col int) bool {
	n := len(grid[0])
	if grid[row][col] == 1 {
		if col == n-1 || grid[row][col+1] == -1 {
			return true
		}
		return false
	}
	if col == 0 || grid[row][col-1] == 1 {
		return true
	}
	return false
}

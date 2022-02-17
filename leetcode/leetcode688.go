package leetcode

//直接模拟法，可能会OOM，空间复杂度8^k
func KnightProbability(n int, k int, row int, column int) float64 {
	if row >= n || column >= n {
		return float64(0)
	}
	if k == 0 {
		return float64(1)
	}

	dx, dy := []int{2, 2, -2, -2, 1, 1, -1, -1}, []int{1, -1, 1, -1, 2, -2, 2, -2}
	res, jumpCount := float64(1), 0
	que := [][]int{{row, column}}
	for jumpCount < k {
		levelSize := len(que)
		if levelSize == 0 {
			return res
		}
		outCount := 0
		for l := 0; l < levelSize; l++ {
			cur := que[0]
			que = que[1:]
			for i := 0; i < 8; i++ {
				nx, ny := cur[0]+dx[i], cur[1]+dy[i]
				if nx < 0 || ny < 0 || nx >= n || ny >= n {
					outCount++
					continue
				}
				que = append(que, []int{nx, ny})
			}
		}
		res *= float64(8*levelSize-outCount) / float64(8*levelSize)
		jumpCount++
	}
	return res
}

//DP做法
func KnightProbability2(n int, k int, row int, column int) float64 {
	dx, dy := []int{2, 2, -2, -2, 1, 1, -1, -1}, []int{1, -1, 1, -1, 2, -2, 2, -2}
	dp := make([][][]float64, k+1)
	for l := 0; l <= k; l++ {
		dp[l] = make([][]float64, n)
		for i := 0; i < n; i++ {
			dp[l][i] = make([]float64, n)
			for j := 0; j < n; j++ {
				if l == 0 {
					dp[l][i][j] = float64(1)
					continue
				}
				for p := 0; p < 8; p++ {
					nx, ny := i+dx[p], j+dy[p]
					if nx < 0 || ny < 0 || nx >= n || ny >= n {
						continue
					}
					dp[l][i][j] += dp[l-1][nx][ny] / float64(8)
				}
			}
		}
	}
	return dp[k][row][column]
}

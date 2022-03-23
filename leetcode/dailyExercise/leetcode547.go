package dailyExercise

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	neighbors := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if isConnected[i][j] == 1 {
				neighbors[i] = append(neighbors[i], j)
				neighbors[j] = append(neighbors[j], i)
			}
		}
	}

	used := make([]bool, n)
	count := 0
	for i := 0; i < n; i++ {
		if !used[i] {
			que := []int{i}
			used[i] = true
			for len(que) > 0 {
				top := que[0]
				que = que[1:]
				for _, neighbor := range neighbors[top] {
					if used[neighbor] {
						continue
					}
					que = append(que, neighbor)
					used[neighbor] = true
				}
			}
			count++
		}
	}
	return count
}

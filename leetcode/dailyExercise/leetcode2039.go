package dailyExercise

func networkBecomesIdle(edges [][]int, patience []int) int {
	n := len(patience)
	distance := calMinDistance(edges, n)
	res := 0
	for i := 0; i < n; i++ {
		doubleDistance := 2 * distance[i]
		if patience[i] >= doubleDistance {
			res = max(res, doubleDistance)
			continue
		}
		divide := doubleDistance / patience[i]
		lastCommit := divide * patience[i]
		if doubleDistance == divide*patience[i] {
			lastCommit = (divide - 1) * patience[i]
		}
		res = max(lastCommit+doubleDistance, res)
	}
	return res + 1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func calMinDistance(edges [][]int, n int) []int {

	distance := make([]int, n)
	used := make([]bool, n)
	neighbors := make(map[int][]int)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1])
		neighbors[edge[1]] = append(neighbors[edge[1]], edge[0])
	}
	que := []int{0}
	used[0] = true
	curDistance := 1
	for len(que) > 0 {
		curLevelSize := len(que)
		for i := 0; i < curLevelSize; i++ {
			top := que[0]
			que = que[1:]
			for _, child := range neighbors[top] {
				if used[child] {
					continue
				}
				distance[child] = curDistance
				que = append(que, child)
				used[child] = true
			}
		}
		curDistance++
	}
	return distance
}

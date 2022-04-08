package dailyExercise

func findMinHeightTrees(n int, edges [][]int) []int {
	neighborsMap := make(map[int][]int)
	for _, edge := range edges {
		neighborsMap[edge[0]] = append(neighborsMap[edge[0]], edge[1])
		neighborsMap[edge[1]] = append(neighborsMap[edge[1]], edge[0])
	}
	resMap := make(map[int][]int)
	min := n
	for i := 0; i < n; i++ {
		tmp := calHeight(i, n, neighborsMap)
		resMap[tmp] = append(resMap[tmp], i)
		if tmp < min {
			min = tmp
		}
	}
	return resMap[min]
}

func calHeight(cur, n int, neighborsMap map[int][]int) int {
	que := []int{cur}
	used := make([]bool, n)
	used[cur] = true
	level := 0
	for len(que) > 0 {
		curLevelSize := len(que)
		for i := 0; i < curLevelSize; i++ {
			top := que[0]
			que = que[1:]
			for _, neigh := range neighborsMap[top] {
				if !used[neigh] {
					que = append(que, neigh)
					used[neigh] = true
				}
			}
		}
		level++
	}
	return level
}

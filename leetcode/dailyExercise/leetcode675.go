package dailyExercise

import "sort"

func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	var trees []int
	distance := make([][]int, m)
	treeHeightToIndexMap := make(map[int][]int)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance[i][j] = -1
			if forest[i][j] > 1 {
				trees = append(trees, forest[i][j])
				treeHeightToIndexMap[forest[i][j]] = []int{i, j}
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i] < trees[j]
	})
	res := 0
	if len(trees) == 0 {
		return 0
	}

	start, _ := calDistance(forest, []int{0, 0}, treeHeightToIndexMap[trees[0]])
	if start == -1 {
		return -1
	}
	res += start
	for i := 0; i < len(trees)-1; i++ {
		curDis, newForest := calDistance(forest, treeHeightToIndexMap[trees[i]], treeHeightToIndexMap[trees[i+1]])
		if curDis == -1 {
			return -1
		}
		res += curDis
		forest = newForest
	}
	return res
}

//不包含起点，包含终点
func calDistance(forest [][]int, from, to []int) (int, [][]int) {
	dx, dy := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}
	m, n := len(forest), len(forest[0])
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	que := [][]int{from}
	level := 0
	used[from[0]][from[1]] = true
	for len(que) > 0 {
		curLevelSize := len(que)
		for i := 0; i < curLevelSize; i++ {
			top := que[0]
			que = que[1:]
			if top[0] == to[0] && top[1] == to[1] {
				return level, forest
			}
			for j := 0; j < 4; j++ {
				nx, ny := top[0]+dx[j], top[1]+dy[j]
				if nx < 0 || nx >= m || ny < 0 || ny >= n || forest[nx][ny] == 0 || used[nx][ny] == true {
					continue
				}
				que = append(que, []int{nx, ny})
				forest[nx][ny] = 1
				used[nx][ny] = true
			}
		}
		level++
	}
	return -1, forest
}

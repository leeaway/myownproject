package dailyExercise

func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	size := make([]int, n)
	//childrenMap: 记录当前节点的孩子节点
	//levelMap: 记录每一层有哪些节点
	childrenMap, levelMap := make(map[int][]int), make(map[int][]int)

	//构造childrenMap
	for i, p := range parents {
		if p == -1 {
			continue
		}
		childrenMap[p] = append(childrenMap[p], i)
	}

	//层序遍历去构造levelMap
	que := []int{0}
	level := 1
	for len(que) > 0 {
		curLevel := len(que)
		for i := 0; i < curLevel; i++ {
			top := que[0]
			levelMap[level] = append(levelMap[level], top)
			que = que[1:]
			for _, node := range childrenMap[top] {
				que = append(que, node)
			}
		}
		level++
	}

	//基于levelMap及size[root] = size[left]+size[right]+1去构造size
	for i := level - 1; i > 0; i-- {
		for _, node := range levelMap[i] {
			size[node] = 1
			for _, c := range childrenMap[node] {
				size[node] += size[c]
			}
		}
	}

	//开始遍历删除每个节点可以得到的分数，并用resToCountMap记录每个结果的个数
	resToCountMap := make(map[int]int)
	max := 0
	for i := 0; i < n; i++ {
		children := childrenMap[i]
		parCount := size[0] - size[i]
		if i == 0 {
			parCount = 1
		}
		for _, c := range children {
			parCount *= size[c]
		}
		resToCountMap[parCount]++
		if max < parCount {
			max = parCount
		}
	}
	return resToCountMap[max]
}

package dailyExercise

func minMutation(start string, end string, bank []string) int {
	que := []string{start}
	used := make(map[string]bool)
	used[start] = true
	res := 0
	for len(que) > 0 {
		curSize := len(que)
		for i := 0; i < curSize; i++ {
			top := que[0]
			if top == end {
				return res
			}
			que = que[1:]
			curNeighbors := getNeighbors(top, bank)
			for _, neigh := range curNeighbors {
				if !used[neigh] {
					que = append(que, neigh)
					used[neigh] = true
				}
			}
		}
		res++
	}
	return -1
}

func getNeighbors(cur string, bank []string) (res []string) {
	for _, b := range bank {
		if canMutation(cur, b) {
			res = append(res, b)
		}
	}
	return res
}

func canMutation(str1, str2 string) bool {
	count := 0
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count++
		}
	}
	return count == 1
}

package dailyExercise

func validUtf8(data []int) bool {
	indexToOneNumMap := make(map[int]int)
	for i, d := range data {
		indexToOneNumMap[i] = getOneNum(d)
		if indexToOneNumMap[i] > 4 {
			return false
		}
	}
	index := 0
	firstOneCount := 0
	for index < len(data) {
		firstOneCount = indexToOneNumMap[index]
		if firstOneCount == 0 {
			index++
			continue
		}
		if firstOneCount == 1 {
			return false
		}
		tmpEnd := index + firstOneCount
		index++
		for index < tmpEnd && index < len(data) {
			if indexToOneNumMap[index] != 1 {
				return false
			}
			index++
		}
		if index != tmpEnd {
			return false
		}
	}
	return true
}

func getOneNum(data int) int {
	start := 7
	for start > 0 && ((data>>start)&1) > 0 {
		start--
	}
	return 7 - start
}

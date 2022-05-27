package dailyExercise

func findClosest(words []string, word1 string, word2 string) int {
	startIndex1, startIndex2 := 300000, 300000
	res := 300000
	for i, word := range words {
		if word == word1 {
			startIndex1 = i
			res = min(res, abs(i-startIndex2))
		}
		if word == word2 {
			startIndex2 = i
			res = min(res, abs(i-startIndex1))
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

package dailyExercise

func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	index, min, max := 1, -1, 1
	for _, c := range s {
		if c == 'I' {
			res[index] = max
			max++
			index++
			continue
		}
		res[index] = min
		min--
		index++
	}
	if min < -1 {
		for i := 0; i <= n; i++ {
			res[i] -= min + 1
		}
	}
	return res
}

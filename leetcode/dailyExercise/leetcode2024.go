package dailyExercise

func maxConsecutiveAnswers(answerKey string, k int) int {
	sum := make([][]int, 2)
	n := len(answerKey)
	sum[0], sum[1] = make([]int, n+1), make([]int, n+1)
	for i, c := range answerKey {
		sum[0][i+1] = sum[0][i]
		sum[1][i+1] = sum[1][i]
		if c == 'T' {
			sum[0][i+1]++
			continue
		}
		sum[1][i+1]++
	}

	return max(maxConsecutiveAnswersChar(answerKey, k, sum[0]), maxConsecutiveAnswersChar(answerKey, k, sum[1]))
}

func maxConsecutiveAnswersChar(answerKey string, k int, sum []int) int {
	l, r, n := 0, 0, len(answerKey)
	res := 0
	for r < n {
		for r < n && sum[r+1]-sum[l] <= k {
			r++
		}
		res = max(res, r-l)
		for r < n && l <= r && sum[r+1]-sum[l] > k {
			l++
		}
	}
	return res
}

package dailyExercise

func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	left, right := make([]int, n), make([]int, n)
	l, r := -1, -1

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1]
		if s[i-1] == '*' {
			sum[i]++
		}
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
		if s[n-1-i] == '|' {
			l = n - 1 - i
		}
		left[n-1-i] = l
	}

	var res []int
	for _, q := range queries {
		x, y := right[q[0]], left[q[1]]
		if x == -1 || y == -1 || x > y {
			res = append(res, 0)
			continue
		}
		res = append(res, sum[y+1]-sum[x+1])
	}
	return res
}

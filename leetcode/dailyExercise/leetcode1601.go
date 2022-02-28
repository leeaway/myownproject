package dailyExercise

//二进制枚举方法
func maximumRequests(n int, requests [][]int) int {
	m := len(requests)
	total := 1 << m
	res := 0
	for i := 0; i < total; i++ {
		var curRequests [][]int
		for j := 0; j < m; j++ {
			if (i & (1 << j)) > 0 {
				curRequests = append(curRequests, requests[j])
			}
		}
		if len(curRequests) <= res {
			continue
		}
		if check(n, curRequests) {
			if len(curRequests) > res {
				res = len(curRequests)
			}
		}
	}
	return res
}

func check(n int, requests [][]int) bool {
	var delta []int
	for i := 0; i < n; i++ {
		delta = append(delta, 0)
	}
	for _, r := range requests {
		delta[r[0]]--
		delta[r[1]]++
	}
	for i := 0; i < n; i++ {
		if delta[i] != 0 {
			return false
		}
	}
	return true
}

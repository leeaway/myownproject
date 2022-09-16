package dailyExercise

func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	total := 0
	for _,matchstick := range matchsticks {
		total += matchstick
	}
	if total%4 > 0 {
		return false
	}
	average := total/4
	for _,matchstick := range matchsticks {
		if matchstick > average {
			return false
		}
	}
	used := make([]bool,n)
	for i:=0;i<3;i++ {
		for j:=0;j<n;j++ {
			if !used[j] {
				used[j] = true
				res := find2(average-matchsticks[j],matchsticks,used)
				if !res {
					return false
				}
			}
		}
	}
	return true
}

func find2(target int,matchsticks []int,used []bool) bool {
	if target == 0 {
		return true
	}
	n := len(matchsticks)
	res := false
	for i:=0;i<n;i++ {
		if used[i] || matchsticks[i] > target {
			continue
		}
		used[i] = true
		res = res || find2(target-matchsticks[i],matchsticks,used)
		used[i] = false
	}
	return res
}





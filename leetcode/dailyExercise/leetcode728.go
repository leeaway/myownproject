package dailyExercise

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for left <= right {
		if isDividingNum(left) {
			res = append(res, left)
		}
		left++
	}
	return res
}

func isDividingNum(n int) bool {
	tmp := n
	degree := 10
	cur := 0
	for tmp > 0 {
		cur = tmp % degree
		if cur == 0 || (n%cur) > 0 {
			return false
		}
		tmp = tmp / degree
	}
	return true
}

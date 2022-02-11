package mathutil

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

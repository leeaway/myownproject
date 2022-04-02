package dailyExercise

import (
	"math"
	"sort"
)

func canReorderDoubled(arr []int) bool {
	var pos, neg, zeros []int
	for _, a := range arr {
		if a < 0 {
			neg = append(neg, a)
			continue
		}
		if a == 0 {
			zeros = append(zeros, a)
			continue
		}
		pos = append(pos, a)
	}
	if ((len(neg) * len(pos)) & 1) > 0 {
		return false
	}
	return helper(pos) && helper(neg) && (len(zeros)&1) == 0
}

func helper(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
	countMap := make(map[int]int)
	for _, a := range arr {
		countMap[a]++
	}
	start := 0
	for start < len(arr) {
		if countMap[arr[start]] == 0 {
			start++
			continue
		}
		if countMap[2*arr[start]] < 1 {
			return false
		}
		countMap[arr[start]]--
		countMap[2*arr[start]]--
		start++
	}
	return true
}

package dailyExercise

import (
	"example.com/m/v2/tools/mathutil"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	vers1 := strings.Split(version1, ".")
	vers2 := strings.Split(version2, ".")

	index := 0
	for index < mathutil.MinInt(len(vers1), len(vers2)) {
		curRes := CompareString(vers1[index], vers2[index])
		if curRes != 0 {
			return curRes
		}
		index++
	}
	for index < len(vers1) {
		if len(strings.TrimLeft(vers1[index], "0")) > 0 {
			return +1
		}
		index++
	}
	for index < len(vers2) {
		if len(strings.TrimLeft(vers2[index], "0")) > 0 {
			return -1
		}
		index++
	}
	return 0
}

func CompareString(s1, s2 string) int {
	s1 = strings.TrimLeft(s1, "0")
	s2 = strings.TrimLeft(s2, "0")
	index := 0
	if len(s1) < len(s2) {
		return -1
	}
	if len(s1) > len(s2) {
		return +1
	}
	for index < len(s2) {
		if s1[index] < s2[index] {
			return -1
		}
		if s1[index] > s2[index] {
			return +1
		}
		index++
	}
	return 0
}

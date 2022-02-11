package leetcode

import "strings"

func compareVersion(version1 string, version2 string) int {
	vers1 := strings.Split(version1, ".")
	vers2 := strings.Split(version2, ".")
	return len(vers2) - len(vers1)
}

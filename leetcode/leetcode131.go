package leetcode

import "fmt"

func Partition(s string) [][]string {
	fmt.Println(dfs(s))
	return nil
}

func dfs(str string) int {
	if len(str) <= 1 {
		return 1
	}
	res := 0
	if IsHuiWen(str) {
		res++
	}
	for i := 1; i <= len(str); i++ {
		res += dfs(str[:i]) * dfs(str[i:])
	}
	return res
}

func IsHuiWen(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

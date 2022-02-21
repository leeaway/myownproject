package leetcode

import "fmt"

func Partition(s string) [][]string {
	fmt.Println(handle(s))
	return nil
}

func handle(str string) [][]string {
	var res [][]string
	if len(str) <= 1 {
		res = append(res, []string{str})
		return res
	}
	if IsHuiWen(str) {
		res = append(res, []string{str})
	}
	for i := 1; i < len(str); i++ {
		if !IsHuiWen(str[:i]) {
			continue
		}
		nextRes := handle(str[i:])
		for j, _ := range nextRes {
			nextRes[j] = append(nextRes[j], str[:i])
		}
		res = append(res, nextRes...)
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

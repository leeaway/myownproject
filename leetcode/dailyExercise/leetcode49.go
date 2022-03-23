package dailyExercise

import "fmt"

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		orderedStr := toOrderedStr(str)
		strMap[orderedStr] = append(strMap[orderedStr], str)
	}
	var res [][]string
	for _, val := range strMap {
		res = append(res, val)
	}
	return res
}

func toOrderedStr(str string) string {
	letterArr := make([]int, 26)
	for _, r := range str {
		letterArr[r-'a']++
	}
	res := ""
	for i := 0; i < 26; i++ {
		if letterArr[i] > 0 {
			res += fmt.Sprintf("%v:%v,", i, letterArr[i])
		}
	}
	return res
}

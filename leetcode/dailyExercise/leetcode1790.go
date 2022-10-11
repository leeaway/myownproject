package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/10/11 10:30
 */

func areAlmostEqual(s1 string, s2 string) bool {
	var diffIndexes []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffIndexes = append(diffIndexes, i)
		}
	}

	if len(diffIndexes) == 0 || (len(diffIndexes) == 2 && s1[diffIndexes[0]] == s2[diffIndexes[1]] && s1[diffIndexes[1]] == s2[diffIndexes[0]]) {
		return true
	}
	return false
}

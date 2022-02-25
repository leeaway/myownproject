package leetcode

func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]bool)
	minWord := 20
	for _, word := range wordDict {
		dictMap[word] = true
		if minWord > len(word) {
			minWord = len(word)
		}
	}
	return solve2(s, dictMap, minWord)
}

func solve2(str string, dictMap map[string]bool, minWord int) bool {
	if len(str) < minWord {
		return false
	}
	if dictMap[str] {
		return true
	}
	res := false
	for i := 0; i < len(str); i++ {
		curStr := str[:i]
		if !dictMap[curStr] {
			continue
		}
		nextRes := solve2(str[i:], dictMap, minWord)
		res = res || nextRes
		if res {
			return true
		}
	}
	return res
}

package leetcode

import "unicode"

func reverseOnlyLetters(s string) string {
	l, r := 0, len(s)-1
	res := []rune(s)
	for l < r {
		for l < r && !unicode.IsLetter(rune(s[l])) {
			l++
		}
		for l < r && !unicode.IsLetter(rune(s[r])) {
			r--
		}
		res[l] = rune(s[r])
		res[r] = rune(s[l])
		l++
		r--
	}
	return string(res)
}

package dailyExercise

import "unicode"

func strongPasswordChecker(password string) int {
	if isValidPwd(password) {
		return 0
	}
	return 1
}

func isValidPwd(pwd string) bool {
	if len(pwd) < 6 || len(pwd) > 20 {
		return false
	}
	hasDigit, hasLowerLetter, hasUpperLetter := false, false, false
	l, r := 0, 0
	var cur rune
	for r < len(pwd) {
		cur = rune(pwd[r])
		if unicode.IsUpper(cur) {
			hasUpperLetter = true
		} else if unicode.IsLower(cur) {
			hasLowerLetter = true
		} else if unicode.IsDigit(cur) {
			hasDigit = true
		}

		for r < len(pwd) && pwd[r] == pwd[l] {
			r++
		}
		if r-l >= 3 && (unicode.IsLetter(cur) || unicode.IsDigit(cur)) {
			return false
		}
		l = r
	}
	return hasDigit && hasLowerLetter && hasUpperLetter
}

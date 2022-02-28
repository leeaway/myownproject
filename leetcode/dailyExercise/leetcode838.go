package dailyExercise

func PushDominoes(dominoes string) string {
	n, l := len(dominoes), 0
	var res []uint8
	for _, d := range dominoes {
		res = append(res, uint8(d))
	}
	leftChar := uint8('L')
	for l < n {
		j := l
		for j < n && dominoes[j] == '.' {
			j++
		}
		rightChar := uint8('R')
		if j < n {
			rightChar = dominoes[j]
		}
		if leftChar == rightChar {
			for l < j {
				res[l] = leftChar
				l++
			}
		} else if leftChar == uint8('R') && rightChar == uint8('L') {
			k := j - 1
			for l < k {
				res[l] = leftChar
				res[k] = rightChar
				l++
				k--
			}
		}
		leftChar = rightChar
		l = j + 1
	}
	return string(res)
}

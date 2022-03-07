package dailyExercise

import (
	"fmt"
	"math"
)

func convertToBase7(num int) string {
	res := ""
	if num < 0 {
		res += "-"
		num = -num
	}
	degree := 1
	for num >= int(math.Pow(float64(7), float64(degree))) {
		degree++
	}

	for degree > 0 {
		degree--
		curNum := int(math.Pow(float64(7), float64(degree)))
		res += fmt.Sprintf("%v", num/curNum)
		num -= curNum * (num / curNum)
	}

	return res
}

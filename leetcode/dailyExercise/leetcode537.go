package dailyExercise

import (
	"fmt"
	"strconv"
	"unicode"
)

func complexNumberMultiply(num1 string, num2 string) string {
	arr1 := splitNum(num1)
	arr2 := splitNum(num2)
	real := arr1[0] * arr2[0]
	imagine := arr1[2] * arr2[2]
	operant := arr1[1] * arr2[1]
	mix := arr2[0]*arr1[2]*arr1[1] + arr1[0]*arr2[2]*arr2[1]
	res := fmt.Sprintf("%v", real-imagine)
	if operant > 0 {
		res += "+"
	} else {
		res += "-"
	}
	if mix < 0 && res[len(res)-1] == '-' {
		return res[:len(res)-1] + "+" + fmt.Sprintf("%vi", -mix)
	}
	return res + fmt.Sprintf("%vi", mix)
}

//分离实部和虚部
func splitNum(num string) []int {
	startOperant := 1
	start := 0
	index := 0
	if !unicode.IsDigit(rune(num[index])) {
		start++
		index++
		startOperant = -1
	}
	for unicode.IsDigit(rune(num[index])) {
		index++
	}
	real, _ := strconv.Atoi(num[start:index])
	operant := 1
	if num[index] == '-' {
		operant = -1
	}
	index++
	imagOperant := 1
	if !unicode.IsDigit(rune(num[index])) {
		imagOperant = -1
		index++
	}
	imagNum, _ := strconv.Atoi(num[index : len(num)-1])
	return []int{real * startOperant, operant, imagNum * imagOperant}
}

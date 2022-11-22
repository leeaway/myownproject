package mathutil

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//最大公约数 greatest common divisor
func Gcd(a, b int) int {
	//if a<b {
	//	a,b = b,a
	//}
	//mod := a%b
	//if mod == 0 {
	//	return b
	//}
	//return Gcd(mod,b)

	//简单写法
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// 最小公倍数： leastCommonMultiple
// a*b = Gcd(a,b) * Lcm(a,b)
func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

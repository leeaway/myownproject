package dailyExercise

var mod int = 1e9+7

func numPrimeArrangements(n int) int {
	primeCount := 0
	for i:=2;i<=n;i++ {
		if isPrime(i) {
			primeCount++
		}
	}

	return factorial(n-primeCount)*factorial(primeCount)%mod


}

func factorial(a int) int {
	res := 1
	for i:=1;i<=a;i++ {
		res = res * i%mod
	}
	return res%mod
}

func isPrime(a int) bool {
	for i:=2;i*i<=a;i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

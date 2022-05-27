package dailyExercise

import "strings"

func findSubstringInWraproundString(p string) int {
	n := len(p)
	occured := make([]bool, 26)
	dp := make([]int, n)
	dp[0] = 1
	occured[p[0]-'a'] = true
	for i := 1; i < n; i++ {
		tmp := i
		if occured[p[i]-'a'] {
			for tmp > 0 && isConn(p[tmp-1], p[tmp]) {
				if !strings.Contains(p[:i], p[tmp-1:i+1]) {
					dp[i]++
				}
				tmp--
			}
		} else {
			dp[i] += 1
			for tmp > 0 && isConn(p[tmp-1], p[tmp]) {
				tmp--
				dp[i]++
			}
		}
		dp[i] += dp[i-1]
		occured[p[i]-'a'] = true
	}
	return dp[n-1]
}

func isConn(a, b uint8) bool {
	return b-a == 1 || (a == 'z' && b == 'a')
}

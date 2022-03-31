package dailyExercise

import (
	"fmt"
	"math"
	"strconv"
)

func kthPalindrome(queries []int, intLength int) []int64 {
	//这里dp[i]记录长度为i的数字包含的回文数数量，包含前导0，如01110也算作5位数
	dp := getCountWithZero(intLength)
	//这里base就是为了修复前导0，比如最后算出来是01110，其实是11111，相当于01110+10001 = 11111，注意这里如果intLength = 1，base应该是1
	base := int64(math.Pow(float64(10), float64(intLength-1))) + 1
	if intLength == 1 {
		base = 1
	}
	var res []int64
	/*
		如果对第一位不包含前导0做特殊逻辑的话，因为只有第一位不能前导0，中间的都可以，后面的find就很难写，很难递归求
		所以直接全包含，我们只需要对最后的逻辑做后处理（因为要从1开始，我们处理是从0开始）：
		1. 如果算出来第一位是9，如919，其实应该是不存在，因为我们是【0，9】，其实应该是【1，9】，所以这里如果算出是9，其实已经超过了，返回-1
		2. 其他情况首末的数字都要加1，如算出来51115，其实应该是61116
	*/
	for _, q := range queries {
		str := find(int64(q), dp, intLength)
		if str[0] == '9' || str == "-1" {
			res = append(res, -1)
			continue
		}

		ans, _ := strconv.ParseInt(str, 10, 64)
		//根据base修正ans
		res = append(res, ans+base)
	}
	return res
}

/*
	这里是一个递归过程，本质上我们每次只要确定最外围的数字即可，比如q=90,l=3,其实是相当于枚举首位数字0-9
	1. 我们假设首位数字为start，则可以排除掉首位<start的所有数，因为从0开始，即排除掉start * dp[l-2],因此每次查找到首位数字后：
		newq = q - start*dp[l-2]
	2. 我们继续求 q = newq, l = l-2的子问题，这样就可以递归了
	3. 递归结束状态：l < 3 或者 q > dp[l](意味着超过所有满足条件的回文数个数)
*/
func find(q int64, dp []int64, l int) string {
	if q > dp[l] {
		return "-1"
	}
	if l == 1 {
		return strconv.FormatInt(q-1, 10)
	}
	if l == 2 {
		return fmt.Sprintf("%v%v", q-1, q-1)
	}
	start := (q - 1) / dp[l-2]
	q -= start * dp[l-2]
	str := strconv.FormatInt(start, 10)
	next := find(q, dp, l-2)
	return str + next + str
}

func getCountWithZero(l int) []int64 {
	dp := make([]int64, 16)
	dp[1] = 10
	dp[2] = 10
	for i := 3; i <= l; i++ {
		dp[i] = 10 * dp[i-2]
	}
	return dp
}

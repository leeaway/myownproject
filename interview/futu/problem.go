package futu

/**
 * @author 2416144794@qq.com
 * @date 2023/1/30 17:04
 */

func checkRepeatedStr(str string) bool {
	n := len(str)
	for i := 1; i <= n/2; i++ {
		if checkHelper(str, str[:i]) {
			return true
		}
	}
	return false
}

func checkHelper(str, sub string) bool {
	subLen, strLen := len(sub), len(str)
	if strLen%subLen > 0 {
		return false
	}
	end := subLen
	for end <= strLen {
		if str[end-subLen:end] != sub {
			return false
		}
		end += subLen
	}
	return true
}

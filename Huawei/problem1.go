package Huawei

import (
	"strconv"
	"strings"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/9 21:28
 */

func solveHelper(str string) int64 {
	var strs []string
	for i := 0; i <= len(str)-2; i += 2 {
		strs = append(strs, str[i:i+2])
	}

	totalDigit := len(strs)
	resStr := ""
	for i, s := range strs {
		cur := Base16ToBase2(s)
		ok, validStr := checkAndReturnValidStr(totalDigit, cur, i)
		if !ok {
			return -1
		}
		resStr += validStr
	}
	res, _ := strconv.ParseInt(resStr, 2, 0)
	return res
}

func checkAndReturnValidStr(totalDigit int, curStr string, curIndex int) (bool, string) {
	//只有一个字节
	if totalDigit == 1 {
		if !strings.HasPrefix(curStr, "0") {
			return false, ""
		}
		return true, curStr[1:]
	}

	//大于1个字节
	prefix := ""
	for i := 0; i < totalDigit; i++ {
		prefix += "1"
	}
	//是第一个
	if curIndex == 0 {
		if !strings.HasPrefix(curStr, prefix) {
			return false, ""
		}
		return true, curStr[len(prefix):]
	}

	if !strings.HasPrefix(curStr, "10") {
		return false, ""
	}
	return true, curStr[2:]
}

func Base16ToBase2(s string) string {
	num, _ := strconv.ParseInt(s, 16, 0)
	//转为二进制
	res := strconv.FormatInt(num, 2)
	l := len(res)
	for i := 0; i < 8-l; i++ {
		res = "0" + res
	}
	return res
}

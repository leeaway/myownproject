package dailyExercise

import "strings"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/7 18:08
 */

//我们有一些二维坐标，如 "(1, 3)" 或 "(2, 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S。返回所有可能的原始字符串到一个列表
//中。
//
// 原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001", "00.01"或一些其他更小的数
//来表示坐标。此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。
//
// 最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格。
//
//
//
//
//示例 1:
//输入: "(123)"
//输出: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
//
//
//
//示例 2:
//输入: "(00011)"
//输出: ["(0.001, 1)", "(0, 0.011)"]
//解释:
//0.0, 00, 0001 或 00.01 是不被允许的。
//
//
//
//示例 3:
//输入: "(0123)"
//输出: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12,
//3)"]
//
//
//
//示例 4:
//输入: "(100)"
//输出: [(10, 0)]
//解释:
//1.0 是不被允许的。
//
//
//
//
// 提示:
//
//
// 4 <= S.length <= 12.
// S[0] = "(", S[S.length - 1] = ")", 且字符串 S 中的其他元素都是数字。
//
//
//
//
// Related Topics 字符串 回溯 👍 115 👎 0

func ambiguousCoordinates(s string) []string {
	numStr := s[1 : len(s)-1]
	var res []string
	for i := 1; i < len(numStr); i++ {
		left := getValidStr(numStr[:i])
		right := getValidStr(numStr[i:])
		for _, l := range left {
			for _, r := range right {
				res = append(res, buildStr(l, r))
			}
		}
	}
	return res
}

func buildStr(l, r string) string {
	return "(" + l + ", " + r + ")"
}

func getValidStr(str string) []string {
	var res []string
	if checkValid(str) {
		res = append(res, str)
	}
	for i := 1; i < len(str); i++ {
		newStr := str[:i] + "." + str[i:]
		if checkValid(newStr) {
			res = append(res, newStr)
		}
	}
	return res
}

func checkValid(str string) bool {
	if strings.Contains(str, ".") {
		nums := strings.Split(str, ".")
		if nums[1][len(nums[1])-1] == '0' {
			return false
		}
		return len(nums[0]) == 1 || nums[0][0] != '0'
	}
	return len(str) == 1 || str[0] != '0'
}

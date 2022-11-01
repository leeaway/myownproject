package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/11/1 20:21
 */

//给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。
//
// 数组表示的字符串 是由数组中的所有元素 按顺序 连接形成的字符串。
//
//
//
// 示例 1：
//
//
//输入：word1 = ["ab", "c"], word2 = ["a", "bc"]
//输出：true
//解释：
//word1 表示的字符串为 "ab" + "c" -> "abc"
//word2 表示的字符串为 "a" + "bc" -> "abc"
//两个字符串相同，返回 true
//
// 示例 2：
//
//
//输入：word1 = ["a", "cb"], word2 = ["ab", "c"]
//输出：false
//
//
// 示例 3：
//
//
//输入：word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
//输出：true
//
//
//
//
// 提示：
//
//
// 1 <= word1.length, word2.length <= 10³
// 1 <= word1[i].length, word2[i].length <= 10³
// 1 <= sum(word1[i].length), sum(word2[i].length) <= 10³
// word1[i] 和 word2[i] 由小写字母组成
//
//
// Related Topics 数组 字符串 👍 73 👎 0

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return mergeStringArray(word1) == mergeStringArray(word2)
}

func mergeStringArray(word []string) string {
	res := ""
	for _, w := range word {
		res += w
	}
	return res
}

//不使用额外空间，双指针
//i,j表示word1[i],word2[j]; p,q表示word1[i][p],word2[j][q]
func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	m, n := len(word1), len(word2)
	i, j, p, q := 0, 0, 0, 0
	for i < m && j < n {
		if word1[i][p] != word2[j][q] {
			return false
		}
		p++
		q++

		if p == len(word1[i]) {
			p = 0
			i++
		}
		if q == len(word2[j]) {
			q = 0
			j++
		}
	}
	return i == m && j == n
}
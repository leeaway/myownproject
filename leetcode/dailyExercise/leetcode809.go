package dailyExercise

import "fmt"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 10:26
 */

//有时候人们会用重复写一些字母来表示额外的感受，比如 "hello" -> "heeellooo", "hi" -> "hiii"。我们将相邻字母都相同的一串
//字符定义为相同字母组，例如："h", "eee", "ll", "ooo"。
//
// 对于一个给定的字符串 S ，如果另一个单词能够通过将一些字母组扩张从而使其和 S 相同，我们将这个单词定义为可扩张的（stretchy）。扩张操作定义如下
//：选择一个字母组（包含字母 c ），然后往其中添加相同的字母 c 使其长度达到 3 或以上。
//
// 例如，以 "hello" 为例，我们可以对字母组 "o" 扩张得到 "hellooo"，但是无法以同样的方法得到 "helloo" 因为字母组 "oo"
//长度小于 3。此外，我们可以进行另一种扩张 "ll" -> "lllll" 以获得 "helllllooo"。如果 S = "helllllooo"，那么查询词
// "hello" 是可扩张的，因为可以对它执行这两种扩张操作使得 query = "hello" -> "hellooo" -> "helllllooo" =
//S。
//
// 输入一组查询单词，输出其中可扩张的单词数量。
//
//
//
// 示例：
//
//
//输入：
//S = "heeellooo"
//words = ["hello", "hi", "helo"]
//输出：1
//解释：
//我们能通过扩张 "hello" 的 "e" 和 "o" 来得到 "heeellooo"。
//我们不能通过扩张 "helo" 来得到 "heeellooo" 因为 "ll" 的长度小于 3 。
//
//
//
//
// 提示：
//
//
// 0 <= len(S) <= 100。
// 0 <= len(words) <= 100。
// 0 <= len(words[i]) <= 100。
// S 和所有在 words 中的单词都只由小写字母组成。
//
//
// Related Topics 数组 双指针 字符串 👍 75 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	思路，将字符串统一成a1b2的形式，如heeelllo表示为h1e3l3o1，然后与对应的word比对
	如：h11e3l3o1 与 h1e1l1o1
*/
func expressiveWords(s string, words []string) int {
	res := 0
	sFormat := buildWordFormat(s)
	for _, word := range words {
		if isExpressive(sFormat, buildWordFormat(word)) {
			res++
		}
	}
	return res
}

//map会覆盖，所以用结构体
type WordItem struct {
	char uint8
	cnt  int
}

func showWordItem(items []*WordItem) string {
	res := ""
	for _, item := range items {
		res += fmt.Sprintf("%c%d", item.char, item.cnt)
	}
	return res
}

//word1:标准
func isExpressive(word1, word2 []*WordItem) bool {
	if len(word1) != len(word2) {
		return false
	}
	for i := 0; i < len(word2); i++ {
		if word1[i].char != word2[i].char || word2[i].cnt > word1[i].cnt {
			return false
		}

		if word2[i].cnt < word1[i].cnt && word1[i].cnt < 3 {
			return false
		}
	}
	return true
}

//如heeelllo表示为h1e3l3o1
func buildWordFormat(str string) []*WordItem {
	var res []*WordItem
	if len(str) == 0 {
		return res
	}
	l, r := 0, 0
	for r < len(str) {
		for r < len(str) && str[r] == str[l] {
			r++
		}
		res = append(res, &WordItem{
			char: str[l],
			cnt:  r - l,
		})
		if r == len(str) {
			return res
		}
		l = r
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

package problem1_9

import "strings"

/**
 * @author 2416144794@qq.com
 * @date 2022/10/11 18:53
 */

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
//
//
//
// 示例 1：
//
// 输入：s = "We are happy."
//输出："We%20are%20happy."
//
//
//
// 限制：
//
// 0 <= s 的长度 <= 10000
//
// Related Topics 字符串 👍 356 👎 0

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

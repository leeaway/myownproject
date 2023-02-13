package slidewindow

/**
 * @author 2416144794@qq.com
 * @date 2023/2/13 17:02
 */

//有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
//
// 假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
//
//
//
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
//
// 你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
//
// 请返回待替换子串的最小可能长度。
//
// 如果原字符串自身就是一个平衡字符串，则返回 0。
//
//
//
// 示例 1：
//
//
//输入：s = "QWER"
//输出：0
//解释：s 已经是平衡的了。
//
// 示例 2：
//
//
//输入：s = "QQWE"
//输出：1
//解释：我们需要把一个 'Q' 替换成 'R'，这样得到的 "RQWE" (或 "QRWE") 是平衡的。
//
//
// 示例 3：
//
//
//输入：s = "QQQW"
//输出：2
//解释：我们可以把前面的 "QQ" 替换成 "ER"。
//
//
// 示例 4：
//
//
//输入：s = "QQQQ"
//输出：3
//解释：我们可以替换后 3 个 'Q'，使 s = "QWER"。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10^5
// s.length 是 4 的倍数
// s 中只含有 'Q', 'W', 'E', 'R' 四种字符
//
//
// Related Topics 字符串 滑动窗口 👍 172 👎 0

/*
	我们维护一个l指针，和r指针，r>l,若s除掉s[l:r]后:
设 partial= n/4，我们选择 s 的一个子串作为「待替换子串」，只有当 s 剩余的部分中 ‘Q’，‘W’，‘E’，‘R’ 的出现次数都小于等于 partial 时，我们才有可能使 s 变为「平衡字符串」。，记录此时的窗口大小，跟最终结果对比；
	找到合法的之后，移动l
	我们用一个map记录数量，map中最多就4个元素，对比是常数级
*/

func balancedString(s string) int {
	charMap := make(map[uint8]int)

	for i, _ := range s {
		charMap[s[i]]++
	}

	if canBalance(charMap, len(s)) {
		return 0
	}

	l, r := 0, 0
	res := len(s)

	for l < len(s) {
		for r < len(s) && !canBalance(charMap, len(s)) {
			charMap[s[r]]--
			r++
		}
		if !canBalance(charMap, len(s)) {
			break
		}
		res = min(res, r-l)
		charMap[s[l]]++
		l++
	}
	return res
}

func canBalance(in map[uint8]int, n int) bool {
	for _, v := range in {
		if v > n/4 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

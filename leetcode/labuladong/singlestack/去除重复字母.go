package singlestack

/**
 * @author 2416144794@qq.com
 * @date 2023/2/10 15:43
 */

//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
//
//
// 示例 1：
//
//
//输入：s = "bcabc"
//输出："abc"
//
//
// 示例 2：
//
//
//输入：s = "cbacdcbc"
//输出："acdb"
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 由小写英文字母组成
//
//
//
//
// 注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-
//distinct-characters 相同
//
// Related Topics 栈 贪心 字符串 单调栈 👍 888 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	先遍历一遍，记录那些字符是重复字符，用duplicateMap记录
	单调栈开始记录：若当前字符小于栈顶&&栈顶为重复字符&&当前索引后还有栈顶元素，出栈
	另外一个stackMap记录单调栈中存在哪些元素，栈中存在的直接continue
*/
func removeDuplicateLetters(s string) string {
	duplicateMap, stackMap, valToLastIndexMap := make(map[uint8]bool), make(map[uint8]bool), make(map[uint8]int)
	counts := [26]uint8{}
	for i, c := range s {
		idx := c - 'a'
		counts[idx]++
		if counts[idx] > 1 {
			duplicateMap[s[i]] = true
		}
		valToLastIndexMap[s[i]] = i
	}

	//内部存索引
	var stack []int

	for i := 0; i < len(s); i++ {
		if stackMap[s[i]] {
			continue
		}

		for len(stack) > 0 && s[stack[len(stack)-1]] > s[i] && duplicateMap[s[stack[len(stack)-1]]] && valToLastIndexMap[s[stack[len(stack)-1]]] > i {
			delete(stackMap, s[stack[len(stack)-1]])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		stackMap[s[i]] = true
	}
	res := ""
	for len(stack) > 0 {
		res += string(s[stack[0]])
		stack = stack[1:]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

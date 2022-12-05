package slidewindow

/**
 * @author 2416144794@qq.com
 * @date 2022/12/1 17:45
 */

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// 1 <= s.length, t.length <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//o(n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 2262 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

/*
	方法一：滑动窗口，实时维护一个map，记录窗口内的字符分布情况，再与target对比，实现ContainsMap(map1,map2)来判断map1是否包含map2
*/
func minWindow(s string, t string) string {
	maps, mapt := make(map[uint8]int), make(map[uint8]int)
	for _, c := range t {
		mapt[uint8(c)]++
	}
	//定义左闭右开窗口 [l,r)
	l, r, n := 0, 0, len(s)
	res := s + t
	for r < n {
		//r一直增加到maps包含mapt为止
		//其实这里不用for，用if也可以，就慢慢靠外层循环进来
		for r < n && !ContainsMap(maps, mapt) {
			maps[s[r]]++
			r++
		}

		//更新left，知道maps不包含mapt为止
		for ContainsMap(maps, mapt) {
			if r-l < len(res) {
				res = s[l:r]
			}
			maps[s[l]]--
			l++
		}
	}
	if len(res) == len(s)+len(t) {
		return ""
	}
	return res
}

func ContainsMap(map1, map2 map[uint8]int) bool {
	for key, val := range map2 {
		if map1[key] < val {
			return false
		}
	}
	return true
}

/*
	方法一其实有可优化点，比如每次都要对比两个map，十分浪费性能，实际上，我们只要计算t中的字符是否都包含即可，只需要实时对比即可
*/

func minWindow2(s string, t string) string {
	maps, mapt := make(map[uint8]int), make(map[uint8]int)
	for _, c := range t {
		mapt[uint8(c)]++
	}
	//定义左闭右开窗口 [l,r)
	l, r, n := 0, 0, len(s)
	//匹配了几个字符，要求数量也匹配
	found := 0
	res := s + t
	for r < n {
		//跟上面一样，其实也可以换成if
		for r < n && found != len(mapt) {
			cur := s[r]
			if mapt[cur] > 0 {
				maps[cur]++
				if maps[cur] == mapt[cur] {
					found++
				}
			}
			r++
		}

		//更新left，知道maps不包含mapt为止
		for found == len(mapt) {
			if r-l < len(res) {
				res = s[l:r]
			}
			needDel := s[l]
			//只操作t中包含的
			if mapt[needDel] > 0 {
				maps[needDel]--
			}
			//判断是否还满足
			if mapt[needDel] > maps[needDel] {
				found--
			}
			l++
		}
	}
	if len(res) == len(s)+len(t) {
		return ""
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

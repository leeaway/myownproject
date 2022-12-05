package slidewindow

/**
 * @author 2416144794@qq.com
 * @date 2022/12/1 18:48
 */

//给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
//
//
// 示例 2：
//
//
//输入：s1= "ab" s2 = "eidboaoo"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
//
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 802 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	mapt := make(map[uint8]int)
	for _, c := range s1 {
		mapt[uint8(c)]++
	}
	//定义左闭右开窗口 [l,r)
	l, r, n := 0, 0, len(s2)
	for r < n {
		for r < n && r-l < len(s1) {
			cur := s2[r]
			if mapt[cur] == 0 {
				l = r + 1
			}
			r++
		}

		for r-l == len(s1) {
			if checkInclusionHelper(s2[l:r], s1) {
				return true
			}
			l++
		}
	}
	return false
}

func checkInclusionHelper(s1, s2 string) bool {
	map1, map2 := make(map[uint8]int), make(map[uint8]int)
	for _, c1 := range s1 {
		map1[uint8(c1)]++
	}
	for _, c2 := range s2 {
		map2[uint8(c2)]++
	}
	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}
	for k, v := range map2 {
		if map1[k] != v {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

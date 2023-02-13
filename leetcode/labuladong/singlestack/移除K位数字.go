package singlestack

/**
 * @author 2416144794@qq.com
 * @date 2023/2/10 15:07
 */

//给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。
//
// 示例 1 ：
//
//
//输入：num = "1432219", k = 3
//输出："1219"
//解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。
//
//
// 示例 2 ：
//
//
//输入：num = "10200", k = 1
//输出："200"
//解释：移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
//
//
// 示例 3 ：
//
//
//输入：num = "10", k = 2
//输出："0"
//解释：从原数字移除所有的数字，剩余为空就是 0 。
//
//
//
//
// 提示：
//
//
// 1 <= k <= num.length <= 10⁵
// num 仅由若干位数字（0 - 9）组成
// 除了 0 本身之外，num 不含任何前导零
//
//
// Related Topics 栈 贪心 字符串 单调栈 👍 916 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	考虑425这个数，移除一个数字，从2开始，因为4>2,所以保留4一定比删除4剩下的数要更大。所以本质上是用单调栈求下一个更小的元素
	如：12345262，移除4个数字，前面栈[1,2,3,4,5],此时2入栈，移除5，4，3，移除了3位，2入栈，此时[1,2,2]
		6入栈[1,2,2,6],最后6出栈，2入栈，刚好移除了4次，最后栈内[1,2,2,2]，即为答案
	特殊情况：
	1. 移除次数已用尽，后面的数直接入栈即可
	2. 移除次数未用尽，按剩余次数依次出栈即可
	3. 处理前导0
*/
func removeKdigits(num string, k int) string {
	var stack []uint8
	for i := 0; i < len(num); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	res := ""
	//去除前导0
	for len(stack) > 0 && stack[0] == '0' {
		stack = stack[1:]
	}
	for len(stack) > 0 {
		res += string(stack[0])
		stack = stack[1:]
	}
	if len(res) == 0 {
		return "0"
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

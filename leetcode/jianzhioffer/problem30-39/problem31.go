package problem30_39

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/24 18:19
 */

//输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈
//的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
//
//
//
// 示例 1：
//
// 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
//输出：true
//解释：我们可以按以下顺序执行：
//push(1), push(2), push(3), push(4), pop() -> 4,
//push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
//
// 示例 2：
//
// 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
//输出：false
//解释：1 不能在 2 之前弹出。
//
//
//
//
// 提示：
//
//
// 0 <= pushed.length == popped.length <= 1000
// 0 <= pushed[i], popped[i] < 1000
// pushed 是 popped 的排列。
//
//
// 注意：本题与主站 946 题相同：https://leetcode-cn.com/problems/validate-stack-sequences/
//
// Related Topics 栈 数组 模拟 👍 391 👎 0

/*
	模拟栈的push和pop
*/
//leetcode submit region begin(Prohibit modification and deletion)
func validateStackSequences(pushed []int, popped []int) bool {
	m, n := len(pushed), len(popped)
	if m != n {
		return false
	}
	if m == 0 {
		return true
	}
	stack := collections.NewStack()
	in, out := 0, 0
	for out < n {
		//若栈为空，直接push
		if stack.IsEmpty() && in < m {
			stack.Push(pushed[in])
			in++
		}

		//若栈顶不等于当前要pop的，继续push
		for in < m && stack.Peek() != popped[out] {
			stack.Push(pushed[in])
			in++
		}

		//栈顶与out相等，pop出来，out递增到下一个
		if stack.Peek() == popped[out] {
			stack.Poll()
			out++
			continue
		}

		//都找不到，return false
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

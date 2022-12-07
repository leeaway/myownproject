package backtrack

/**
 * @author 2416144794@qq.com
 * @date 2022/12/5 20:55
 */

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 字符串 动态规划 回溯 👍 2987 👎 0

/*
方法一：
括号有效性：
1.有效的括号组合一定是以右括号结尾；
2.截止到任一索引，左括号的数量一定不小于右括号的数量

因此，我们定义l,r为左右括号的剩余数量，采用回溯的方法做：
	则有：r<l(右括号剩余的少，则说明用的更多)，直接return
*/

func generateParenthesis(n int) []string {
	return generateParenthesisHelper(n, n, "")
}

func generateParenthesisHelper(l, r int, cur string) []string {
	var res []string
	//递归终止条件
	if l < 0 || r < 0 || r < l {
		return res
	}
	//都用完了，加入结果集
	if l == 0 && r == 0 {
		return append(res, cur)
	}
	//加入左括号
	res = append(res, generateParenthesisHelper(l-1, r, cur+"(")...)
	res = append(res, generateParenthesisHelper(l, r-1, cur+")")...)
	return res
}

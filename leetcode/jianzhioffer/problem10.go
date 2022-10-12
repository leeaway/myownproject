package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/10/12 18:33
 */

//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//
//F(0) = 0,   F(1) = 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//
// 斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//
//
//
// 示例 1：
//
//
//输入：n = 2
//输出：1
//
//
// 示例 2：
//
//
//输入：n = 5
//输出：5
//
//
//
//
// 提示：
//
//
// 0 <= n <= 100
//
//
// Related Topics 记忆化搜索 数学 动态规划 👍 417 👎 0

/*
方法一：动态规划
*/
func fib_dp(n int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return n
	}
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}
	return dp[n]
}

/*
方法二：递归，效率很低
*/
func fib_recursive(n int) int {
	if n < 2 {
		return n
	}
	return (fib_recursive(n-1) + fib_recursive(n-2)) % (1e9 + 7)
}

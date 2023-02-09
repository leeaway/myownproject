package labuladong

/**
 * @author 2416144794@qq.com
 * @date 2023/1/29 14:56
 */

//0,1,···,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字（删除后从下一个数字开始计数）。求出这个圆圈里剩下的最后一个数字。
//
//
// 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
//
//
//
// 示例 1：
//
//
//输入: n = 5, m = 3
//输出: 3
//
//
// 示例 2：
//
//
//输入: n = 10, m = 17
//输出: 2
//
//
//
//
// 限制：
//
//
// 1 <= n <= 10^5
// 1 <= m <= 10^6
//
//
// Related Topics 递归 数学 👍 722 👎 0

/*
	定义dp(n,m)为在n个数字序列（0~n-1）中不断删除第m个（注意这里是从1开始数）数字最后的结果
	dp[5,3] 初始数组：[0,1,2,3,4]
	dp[4,3] 初始数组：[0,1,2,3]
	dp[5,3] 删除一个数后，dp[5-1,3]对应数组：[3,4,0,1],与dp[4,3]的关系是：
	(0+m)%n = 3%5 = 3
	(1+m)%n = 4%5 = 4
	(2+m)%n = 5%5 = 0
	(3+m)%n = 6%5 = 1

=> dp[n,m] = (dp[n-1,m]+m) % n
dp[n,m]只跟dp[n-1,m]有关，所以可以用一个变量来代替
*/

func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i < n; i++ {
		ans = (ans + m) % n
	}
	return ans
}

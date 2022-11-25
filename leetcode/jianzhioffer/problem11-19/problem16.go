package problem11_19

/**
 * @author 2416144794@qq.com
 * @date 2022/10/13 22:45
 */

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xⁿ）。不得使用库函数，同时不需要考虑大数问题。
//
//
//
// 示例 1：
//
//
//输入：x = 2.00000, n = 10
//输出：1024.00000
//
//
// 示例 2：
//
//
//输入：x = 2.10000, n = 3
//输出：9.26100
//
// 示例 3：
//
//
//输入：x = 2.00000, n = -2
//输出：0.25000
//解释：2⁻² = 1/2² = 1/4 = 0.25
//
//
//
// 提示：
//
//
// -100.0 < x < 100.0
// -2³¹ <= n <= 2³¹-1
// -10⁴ <= xⁿ <= 10⁴
//
//
//
//
// 注意：本题与主站 50 题相同：https://leetcode-cn.com/problems/powx-n/
//
// Related Topics 递归 数学 👍 357 👎 0

/*
方法：快速幂，即x^n = x^a*x^b,其中a+b=n;
	显然a与b越接近效率越高，即：
	n%2==0: a=b=n/2
	n%2!=0: a=n/2,b=n-a
*/

//递归做法
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0.0
	}
	if n < 0 {
		return 1.0 / cal(x, -n)
	}
	return cal(x, n)
}

func cal(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	y := cal(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

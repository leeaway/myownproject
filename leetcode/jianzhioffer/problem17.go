package jianzhioffer

import "math"

/**
 * @author 2416144794@qq.com
 * @date 2022/10/13 23:00
 */

//输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
//
// 示例 1:
//
// 输入: n = 1
//输出: [1,2,3,4,5,6,7,8,9]
//
//
//
//
// 说明：
//
//
// 用返回一个整数列表来代替打印
// n 为正整数
//
//
// Related Topics 数组 数学 👍 259 👎 0

func printNumbers(n int) []int {
	max := int(math.Pow(10.0, float64(n)))
	res := make([]int, max-1)
	cur := 1
	for cur < max {
		res[cur-1] = cur
		cur++
	}
	return res
}

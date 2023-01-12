package problem60_69

/**
 * @author 2416144794@qq.com
 * @date 2023/1/3 18:41
 */

//给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积, 即 B[
//i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
//
//
//
// 示例:
//
//
//输入: [1,2,3,4,5]
//输出: [120,60,40,30,24]
//
//
//
// 提示：
//
//
// 所有元素乘积之和不会溢出 32 位整数
// a.length <= 100000
//
//
// Related Topics 数组 前缀和 👍 282 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	题目即求A[0,i-1]*A[i+1,n-1],其中n = len(nums),A[0,i-1]表示nums[0]到nums[i-1]的乘积
	A[0,i-1]好求，只需要跟前缀和一样求就可以；
	A[i+1,n-1]怎么求，其实反过来，从数组末尾类似前缀和求即可
*/

func constructArr(a []int) []int {
	n := len(a)
	first, last := make([]int, n+1), make([]int, n+1)
	first[0], last[n] = 1, 1
	for i := 0; i < n; i++ {
		first[i+1] = first[i] * a[i]
		last[n-i-1] = last[n-i] * a[n-i-1]
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = first[i] * last[i+1]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

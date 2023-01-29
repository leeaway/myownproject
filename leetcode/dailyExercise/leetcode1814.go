package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2023/1/17 11:38
 */

//给你一个数组 nums ，数组中只包含非负整数。定义 rev(x) 的值为将整数 x 各个数字位反转得到的结果。比方说 rev(123) = 321 ，
//rev(120) = 21 。我们称满足下面条件的下标对 (i, j) 是 好的 ：
//
//
// 0 <= i < j < nums.length
// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
//
//
// 请你返回好下标对的数目。由于结果可能会很大，请将结果对 10⁹ + 7 取余 后返回。
//
//
//
// 示例 1：
//
// 输入：nums = [42,11,1,97]
//输出：2
//解释：两个坐标对为：
// - (0,3)：42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121 。
// - (1,2)：11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12 。
//
//
// 示例 2：
//
// 输入：nums = [13,10,35,24,76]
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁹
//
//
// Related Topics 数组 哈希表 数学 计数 👍 59 👎 0

/*
	nums[i] + rev(nums[j]) == nums[j] + rev(nums[i]) 可得：
		nums[i] - rev(nums[i])  == nums[j] - rev(nums[j])
1. 我们用一个新的数组 ans记录 nums[i] - rev(nums[i])的值，即ans[i] = nums[i] - rev(nums[i])
题目就变成了求ans中满足ans[i] = ans[j]的个数，这个用Hash表就可以计数
2. 计数本质上就是C(n,2) = n*(n-1)/2
*/
//leetcode submit region begin(Prohibit modification and deletion)
func countNicePairs(nums []int) int {
	n := len(nums)
	mod := 1000000007
	ans := make([]int, n)
	cntMap := make(map[int]int)
	res := 0
	for i := 0; i < n; i++ {
		ans[i] = nums[i] - recv(nums[i])
		res = (res + cntMap[ans[i]]) % mod
		cntMap[ans[i]]++
	}
	return res
}

func recv(num int) int {
	res := 0
	for num > 0 {
		cur := num % 10
		num = num / 10
		res = res*10 + cur
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

package problem110_119

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/29 11:32
 */

//给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
//
//
// 示例 1：
//
//
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
//
// 示例 2：
//
//
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
//
//
//
//
// 进阶：可以设计并实现时间复杂度为 O(n) 的解决方案吗？
//
//
//
//
// 注意：本题与主站 128 题相同： https://leetcode-cn.com/problems/longest-consecutive-
//sequence/
//
// Related Topics 并查集 数组 哈希表 👍 56 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
	传统方法排个序，然后维护一个滑动窗口即可，时间复杂度是O(nlog(n))
	实现O(n)需要用并查集
*/
func longestConsecutive(nums []int) int {
	uf := collections.NewUnionFindByMap(nums)
	existMap := make(map[int]bool)
	for _, num := range nums {
		existMap[num] = true
		if existMap[num-1] {
			uf.Union(num, num-1)
		}
		if existMap[num+1] {
			uf.Union(num, num+1)
		}
	}
	return uf.GetMaxHerdSize()
}

//leetcode submit region end(Prohibit modification and deletion)

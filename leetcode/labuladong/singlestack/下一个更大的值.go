package singlestack

/**
 * @author 2416144794@qq.com
 * @date 2022/12/29 18:09
 */

//nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
//
// 给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
//
// 对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定
//nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
//
// 返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出：[-1,3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
//- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
//- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
//
// 示例 2：
//
//
//输入：nums1 = [2,4], nums2 = [1,2,3,4].
//输出：[3,-1]
//解释：nums1 中每个值的下一个更大元素如下所述：
//- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
//- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length <= nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 10⁴
// nums1和nums2中所有整数 互不相同
// nums1 中的所有整数同样出现在 nums2 中
//
//
//
//
// 进阶：你可以设计一个时间复杂度为 O(nums1.length + nums2.length) 的解决方案吗？
//
// Related Topics 栈 数组 哈希表 单调栈 👍 904 👎 0

/*
	单调栈通常用于解决下一个更大或更小值
*/

//leetcode submit region begin(Prohibit modification and deletion)
//我们先处理nums2，记录nums2中下一个更大的数
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//栈内存nums2的数
	var stack []int
	//补一个后位0方便处理
	nums2 = append(nums2, 0)
	//记录结果
	resMap := make(map[int]int)
	for _, num := range nums2 {
		//栈非空且栈顶小于当前元素，直接pop，并记录结果集
		for len(stack) > 0 && stack[len(stack)-1] < num {
			resMap[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	for len(stack) > 0 {
		resMap[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	var res []int
	for _, num := range nums1 {
		res = append(res, resMap[num])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//下一个更小的数
func nextLessElement(nums []int) []int {
	res := make([]int, len(nums))
	//单调递减栈,栈内存索引
	var stack []int

	for i, num := range nums {
		//非空且栈顶元素大于当前数
		for len(stack) > 0 && nums[stack[len(stack)-1]] > num {
			res[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		res[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}
	return res
}

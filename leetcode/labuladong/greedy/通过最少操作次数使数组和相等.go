package greedy

import "sort"

/**
 * @author 2416144794@qq.com
 * @date 2022/12/7 11:34
 */

//给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）。
//
// 每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）。
//
// 请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1 。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
//输出：3
//解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
//- 将 nums2[0] 变为 6 。 nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2] 。
//- 将 nums1[5] 变为 1 。 nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2] 。
//- 将 nums1[2] 变为 2 。 nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2] 。
//
//
// 示例 2：
//
// 输入：nums1 = [1,1,1,1,1,1,1], nums2 = [6]
//输出：-1
//解释：没有办法减少 nums1 的和或者增加 nums2 的和使二者相等。
//
//
// 示例 3：
//
// 输入：nums1 = [6,6], nums2 = [1]
//输出：3
//解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
//- 将 nums1[0] 变为 2 。 nums1 = [2,6], nums2 = [1] 。
//- 将 nums1[1] 变为 2 。 nums1 = [2,2], nums2 = [1] 。
//- 将 nums2[0] 变为 4 。 nums1 = [2,2], nums2 = [4] 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// 1 <= nums1[i], nums2[i] <= 6
//
//
// Related Topics 贪心 数组 哈希表 计数 👍 80 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := getSum(nums1), getSum(nums2)
	if sum1 == sum2 {
		return 0
	}
	diff := sum1 - sum2
	if sum1 < sum2 {
		diff = sum2 - sum1
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}

	/*
		上面已经保证sum1>sum2,且diff为正数，则将nums1降序排序,nums2升序排序
		然后双指针，判断先操作哪个数组可以更大的缩小diff，就移动哪个指针，这里注意特殊情况：
		若两个数组都不能缩小diff，则直接判断当前diff是否大于0即可，否则会无限循环；
	*/

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] > nums1[j]
	})

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	cnt := 0
	m, n := len(nums1), len(nums2)
	p1, p2 := 0, 0
	for diff > 0 && (p1 < m || p2 < n) {
		canMinus1, canMinus2 := 0, 0
		if p1 < m {
			canMinus1 = nums1[p1] - 1
		}
		if p2 < n {
			canMinus2 = 6 - nums2[p2]
		}

		if canMinus1 == 0 && canMinus2 == 0 && diff > 0 {
			return -1
		}

		if canMinus1 > canMinus2 {
			diff -= canMinus1
			p1++
		} else {
			diff -= canMinus2
			p2++
		}
		cnt++
	}
	if diff > 0 {
		return -1
	}
	return cnt
}

func getSum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

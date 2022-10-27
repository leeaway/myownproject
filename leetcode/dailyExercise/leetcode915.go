package dailyExercise

/**
 * @author 2416144794@qq.com
 * @date 2022/10/24 11:43
 */

//给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：
//
//
// left 中的每个元素都小于或等于 right 中的每个元素。
// left 和 right 都是非空的。
// left 的长度要尽可能小。
//
//
// 在完成这样的分组后返回 left 的 长度 。
//
// 用例可以保证存在这样的划分方法。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁶
// 可以保证至少有一种方法能够按题目所描述的那样对 nums 进行划分。
//
//
// Related Topics 数组 👍 153 👎 0

func partitionDisjoint(nums []int) int {
	n := len(nums)
	minV, maxV := make([]int, n), make([]int, n)
	l, r := 0, n-1

	for r >= 0 {
		if r == n-1 {
			minV[r] = nums[r]
		} else {
			minV[r] = min(minV[r+1], nums[r])
		}
		r--
	}

	for l < n-1 {
		if l == 0 {
			maxV[l] = nums[l]
		} else {
			maxV[l] = max(maxV[l-1], nums[l])
		}
		if maxV[l] <= minV[l+1] {
			return l + 1
		}
		l++
	}

	return n
}

func partitionDisjoint2(nums []int) int {
	leftMax := nums[0]
	index, maxN := 0, nums[0]
	for i, num := range nums {
		if num < leftMax {
			leftMax = maxN
			index = i
		} else {
			maxN = max(num, maxN)
		}
	}
	return index + 1
}

package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/10/11 16:38
 */

//找出数组中重复的数字。
//
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
//请找出数组中任意一个重复的数字。
//
// 示例 1：
//
// 输入：
//[2, 3, 1, 0, 2, 5, 3]
//输出：2 或 3
//
//
//
//
// 限制：
//
// 2 <= n <= 100000
//
// Related Topics 数组 哈希表 排序 👍 971 👎 0

/*
方法一：哈希表法 思路： 遍历数组，使用一个 map[int]bool 来记录每一个数是否出现过，遇到已经 出现过的数直接返回。
时间复杂度：O(n) n 是数组元素个数，最坏情况下我们需要遍历完整个数组。
空间复杂度：O(n) n 是数组元素个数，最坏情况下我们需要记录所有数字是否已经出现过。
*/

func findRepeatNumber1(nums []int) int {
	resMap := make(map[int]bool)
	for _, num := range nums {
		if resMap[num] {
			return num
		}
		resMap[num] = true
	}
	return -1
}

/*
	方法二：数组重排序 思路： 已知长度为 n 的数组 nums 里面的所有数字都在 0~n-1 的范围内，因此当 我们对数组元素进行排序后，如果数组中没有重复元素，则排序后的数组，其 每一位 nums[i] == i，如果有重复元素，那么必然会有多个位置出现相同的 元素，而有些位置则没有与该下标相同的元素，由此我们可以通过对数组排序 来找到重复元素。
	我们遍历原数组，对每一个 nums[i] 做如下处理：
	1、判断 nums[i] == i ?
		ture： i++
		false： 判断 nums[i] == nums[nums[i]] ?
			true： 遇到重复元素，返回 nums[i]
			false: 交换 nums[i] 与 nums[nums[i]]，把 nums[i] 的 值交换到它排序后该出现的位置。
	2、重复 1 的过程，直至找到重复元素或是遍历结束
	时间复杂度：O(n) 我们需要遍历数组进行交换排序，对于每一个位置我们最多交换两次，故时间 复杂度为 O(n)。
	空间复杂度：O(1)
*/

func findRepeatNumber2(nums []int) int {
	n := len(nums)
	index := 0
	for index < n {
		if nums[index] == index {
			index++
			continue
		}
		if nums[nums[index]] == nums[index] {
			return nums[index]
		}
		nums[index], nums[nums[index]] = nums[nums[index]], nums[index]
	}
	return -1
}

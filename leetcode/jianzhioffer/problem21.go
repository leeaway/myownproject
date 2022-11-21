package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/11/21 19:46
 */

//输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
//
//
//
// 示例：
//
//
//输入：nums = [1,2,3,4]
//输出：[1,3,2,4]
//注：[3,1,2,4] 也是正确的答案之一。
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 50000
// 0 <= nums[i] <= 10000
//
//
// Related Topics 数组 双指针 排序 👍 268 👎 0

//方法一：首尾双指针
//定义头指针 left ，尾指针 right .
//left 一直往右移，直到它指向的值为偶数
//right 一直往左移， 直到它指向的值为奇数
//交换 nums[left] 和 nums[right] .
//重复上述操作，直到 left==right .
func exchange(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	l, r := 0, n-1
	for l < r {
		for l < r && nums[l]%2 > 0 {
			l++
		}
		for r > l && nums[r]%2 == 0 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}

//快慢指针，slow表示下一个奇数应该在的位置，fast向前搜索奇数，找到就跟slow交换
func exchange2(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast]%2 > 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}

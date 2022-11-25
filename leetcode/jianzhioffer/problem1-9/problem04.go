package problem1_9

/**
 * @author 2416144794@qq.com
 * @date 2022/10/11 17:08
 */

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个
//整数，判断数组中是否含有该整数。
//
//
//
// 示例:
//
// 现有矩阵 matrix 如下：
//
//
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
//
//
// 给定 target = 5，返回 true。
//
// 给定 target = 20，返回 false。
//
//
//
// 限制：
//
// 0 <= n <= 1000
//
// 0 <= m <= 1000
//
//
//
// 注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
//
// Related Topics 数组 二分查找 分治 矩阵 👍 805 👎 0

/*
方法一：
	1.从左下角开始，每次对比可排除一行或一列
	时间复杂度：O(m+n)
	空间复杂度：O(1)
*/
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	row, col := m-1, 0
	for row >= 0 && col < n {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			row--
		} else {
			col++
		}
	}
	return false
}

/*
方法二：
	1.按行或列二分
	时间复杂度：O(m*log(n))
	空间复杂度：O(1)
	只有再n比m大很多的时候，才会比方法一高效
*/
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	//n := len(matrix[0])
	for i := 0; i < m; i++ {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r + 1) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			r = mid - 1
		}
		if nums[mid] < target {
			l = mid + 1
		}
	}
	return false
}

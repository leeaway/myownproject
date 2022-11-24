package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/11/23 20:47
 */

//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
//
//
//
// 示例 1：
//
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//
// 示例 2：
//
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//
//
//
// 限制：
//
//
// 0 <= matrix.length <= 100
// 0 <= matrix[i].length <= 100
//
//
// 注意：本题与主站 54 题相同：https://leetcode-cn.com/problems/spiral-matrix/
//
// Related Topics 数组 矩阵 模拟 👍 470 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/*
 模拟即可，维护四个方向的边界，按照向右->向下->向左->向上->向右的顺序循环，直到上边界 = 下边界 && 左边界 = 右边界
*/
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	//边界
	d, u, l, r := m-1, 0, 0, n-1
	//方向
	down, up, left, right := 0, 1, 2, 3
	curDir := left
	var res []int
	cur := 0
	for len(res) < m*n {
		if curDir == left && len(res) < m*n {
			for i := l; i <= r; i++ {
				res = append(res, matrix[u][i])
				cur++
			}
			u++
			curDir = down
		}
		if curDir == down && len(res) < m*n {
			for i := u; i <= d; i++ {
				res = append(res, matrix[i][r])
				cur++
			}
			r--
			curDir = right
		}
		if curDir == right && len(res) < m*n {
			for i := r; i >= l; i-- {
				res = append(res, matrix[d][i])
				cur++
			}
			d--
			curDir = up
		}
		if curDir == up && len(res) < m*n {
			for i := d; i >= u; i-- {
				res = append(res, matrix[i][l])
				cur++
			}
			l++
			curDir = left
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

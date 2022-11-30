package problem30_39

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 18:18
 */

//从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
//
//
//
// 例如: 给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
//
//
//
//
// 提示：
//
//
// 节点总数 <= 1000
//
//
// 注意：本题与主站 102 题相同：https://leetcode-cn.com/problems/binary-tree-level-order-
//traversal/
//
// Related Topics 树 广度优先搜索 二叉树 👍 256 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder1(root *collections.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	que := []*collections.TreeNode{root}
	level := 0
	for len(que) > 0 {
		levelSize := len(que)
		res = append(res, []int{})
		for i := 0; i < levelSize; i++ {
			top := que[0]
			res[level] = append(res[level], top.Val)
			que = que[1:]
			if top.Left != nil {
				que = append(que, top.Left)
			}
			if top.Right != nil {
				que = append(que, top.Right)
			}
		}
		level++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

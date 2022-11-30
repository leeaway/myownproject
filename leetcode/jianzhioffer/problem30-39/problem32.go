package problem30_39

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 15:10
 */

//从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
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
// 返回：
//
// [3,9,20,15,7]
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
// Related Topics 树 广度优先搜索 二叉树 👍 237 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
	典型的广度优先搜索，实现层序遍历
*/
func levelOrder(root *collections.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	que := []*collections.TreeNode{root}
	res = append(res, root.Val)
	for len(que) > 0 {
		top := que[0]
		que = que[1:]
		if top.Left != nil {
			que = append(que, top.Left)
			res = append(res, top.Left.Val)
		}
		if top.Right != nil {
			que = append(que, top.Right)
			res = append(res, top.Right.Val)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

package binarytree

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2023/1/31 18:46
 */

//给定一个二叉树 根节点 root ，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
//
// 节点 node 的子树为 node 本身，以及所有 node 的后代。
//
//
//
// 示例 1:
//
//
//输入: [1,null,0,0,1]
//输出: [1,null,0,null,1]
//解释:
//只有红色节点满足条件“所有不包含 1 的子树”。
//右图为返回的答案。
//
//
//
//
// 示例 2:
//
//
//输入: [1,0,1,0,0,0,1]
//输出: [1,null,1,null,1]
//解释:
//
//
//
//
// 示例 3:
//
//
//输入: [1,1,0,1,1,0,1,0]
//输出: [1,1,0,1,1,null,1]
//解释:
//
//
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,200]
// 二叉树节点的值只会是 0 或 1
//
//
//
//
//
// 注意：本题与主站 814 题相同：https://leetcode-cn.com/problems/binary-tree-pruning/
//
// Related Topics 树 深度优先搜索 二叉树 👍 60 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *collections.TreeNode) *collections.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}

func isZeroTree(node *collections.TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Val > 0 {
		return false
	}
	if node.Left != nil && node.Left.Val > 0 {
		return false
	}
	if node.Right != nil && node.Right.Val > 0 {
		return false
	}
	return isZeroTree(node.Left) && isZeroTree(node.Right)
}

//leetcode submit region end(Prohibit modification and deletion)

package binarytree

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2023/2/1 15:37
 */

//给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
//输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//
//
// 示例 2：
//
//
//
//
//输入：root = [5,1,7]
//输出：[1,null,5,null,7]
//
//
//
//
// 提示：
//
//
// 树中节点数的取值范围是 [1, 100]
// 0 <= Node.val <= 1000
//
//
//
//
//
// 注意：本题与主站 897 题相同： https://leetcode-cn.com/problems/increasing-order-search-
//tree/
//
// Related Topics 栈 树 深度优先搜索 二叉搜索树 二叉树 👍 50 👎 0

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
	先中序遍历，后重建
*/

var resList []int

func increasingBST(root *collections.TreeNode) *collections.TreeNode {
	inorder(root)
	ans := collections.NewTreeNode(-1)
	dummy := ans
	for _, r := range resList {
		dummy.Right = collections.NewTreeNode(r)
		dummy = dummy.Right
	}
	return ans.Right
}

func inorder(root *collections.TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	resList = append(resList, root.Val)
	inorder(root.Right)
}

/*
	方法二，在中序遍历的过程中改变指向，只需要将上一个遍历到的节点的右孩子指向当前，当前的左孩子指向nil即可
*/

var lastNode *collections.TreeNode

func increasingBST2(root *collections.TreeNode) *collections.TreeNode {
	dummy := collections.NewTreeNode(-1)
	lastNode = dummy
	inorderHelper(lastNode)
	return dummy.Right
}

func inorderHelper(root *collections.TreeNode) {
	if root == nil {
		return
	}
	inorderHelper(root.Left)
	lastNode.Right = root
	root.Left = nil
	lastNode = root
	inorderHelper(root.Right)
}

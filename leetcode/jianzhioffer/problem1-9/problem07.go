package problem1_9

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/10/12 11:00
 */

//输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
//
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
//
//
//
// 示例 1:
//
//
//Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//Output: [3,9,20,null,null,15,7]
//
//
// 示例 2:
//
//
//Input: preorder = [-1], inorder = [-1]
//Output: [-1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 5000
//
//
//
// 注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-
//preorder-and-inorder-traversal/
//
// Related Topics 树 数组 哈希表 分治 二叉树 👍 917 👎 0

/*
方法：preorder中第一个记为根节点，然后找到inorder中的根节点的位置，左边记为左子树，右边即为右子树
	如：Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	此时左子树为9，右子树为15,20,7，根据长度去preorder中切分出左子树的preorder为[9],右子树的preorder为[20,15,7]，递归处理即可；
	递归终止条件：
		len(pre) == 0
*/

func buildTree(preorder []int, inorder []int) *collections.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &collections.TreeNode{
		Val: preorder[0],
	}
	inMap := valToIndexMap(inorder)
	rootIndexInOrder := inMap[preorder[0]]
	root.Left = buildTree(preorder[1:rootIndexInOrder+1], inorder[0:rootIndexInOrder])
	if rootIndexInOrder < len(inorder)-1 {
		rightTreeNodeCount := len(inorder) - rootIndexInOrder - 1
		root.Right = buildTree(preorder[len(preorder)-rightTreeNodeCount:], inorder[rootIndexInOrder+1:])
	}
	return root
}

func valToIndexMap(nums []int) map[int]int {
	res := make(map[int]int)
	for i, num := range nums {
		res[num] = i
	}
	return res
}

func preOrder(root *collections.TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = preOrder(root.Left, res)
	res = preOrder(root.Right, res)
	return res
}

func inOrder(root *collections.TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = inOrder(root.Left, res)
	res = append(res, root.Val)
	res = inOrder(root.Right, res)
	return res
}

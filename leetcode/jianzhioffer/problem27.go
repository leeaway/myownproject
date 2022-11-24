package jianzhioffer

/**
 * @author 2416144794@qq.com
 * @date 2022/11/23 20:30
 */

//请完成一个函数，输入一个二叉树，该函数输出它的镜像。
//
// 例如输入：
//
// 4 / \ 2 7 / \ / \ 1 3 6 9 镜像输出：
//
// 4 / \ 7 2 / \ / \ 9 6 3 1
//
//
//
// 示例 1：
//
// 输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 1000
//
// 注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 321 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 本质上就是左右子树互换，注意要先保存副本
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Right
	root.Right = mirrorTree(root.Left)
	root.Left = mirrorTree(tmp)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

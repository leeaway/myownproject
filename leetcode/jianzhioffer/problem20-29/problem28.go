package problem20_29

/**
 * @author 2416144794@qq.com
 * @date 2022/11/23 20:36
 */

//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
// 1 / \ 2 2 / \ / \ 3 4 4 3 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
// 1 / \ 2 2 \ \ 3 3
//
//
//
// 示例 1：
//
// 输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
// 输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
// 限制：
//
// 0 <= 节点个数 <= 1000
//
// 注意：本题与主站 101 题相同：https://leetcode-cn.com/problems/symmetric-tree/
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 398 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(A, B *TreeNode) bool {
	if A == nil || B == nil {
		return A == nil && B == nil
	}
	return A.Val == B.Val && checkSymmetric(A.Left, B.Right) && checkSymmetric(A.Right, B.Left)
}

//leetcode submit region end(Prohibit modification and deletion)

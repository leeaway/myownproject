package problem30_39

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/27 17:22
 */

//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//
//
// 示例 2：
//
//
//
//
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1,2], targetSum = 0
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点总数在范围 [0, 5000] 内
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
// 注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/
//
// Related Topics 树 深度优先搜索 回溯 二叉树 👍 381 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type collections.TreeNode struct {
 *     Val int
 *     Left *collections.TreeNode
 *     Right *collections.TreeNode
 * }
 */

/*
	典型的dfs遍历所有路径，将路径和等于target的保留下来即可
*/

func pathSum(root *collections.TreeNode, target int) (res [][]int) {
	var path []int
	var dfs func(node *collections.TreeNode, left int)
	dfs = func(node *collections.TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.IsLeaf() && left == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, target)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

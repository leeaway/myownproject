package problem30_39

import "example.com/m/v2/tools/collections"

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 18:42
 */

//请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
//  [20,9],
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
// Related Topics 树 广度优先搜索 二叉树 👍 259 👎 0

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
	同样的层序遍历，只是添加到结果集的方向不一样：
	1.奇数层为从左到右，每次都加在尾部；
	2.偶数层为从右到左，每次都加在头部；
*/
func levelOrder2(root *collections.TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	que := []*collections.TreeNode{root}
	//注意根节点为第0层,因此奇偶对换
	level := 0
	for len(que) > 0 {
		levelSize := len(que)
		res = append(res, []int{})
		for i := 0; i < levelSize; i++ {
			top := que[0]
			//偶数level从左到右
			if level%2 == 0 {
				res[level] = append(res[level], top.Val)
			} else {
				res[level] = append([]int{top.Val}, res[level]...)
			}
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

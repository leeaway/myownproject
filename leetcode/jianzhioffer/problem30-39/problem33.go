package problem30_39

/**
 * @author 2416144794@qq.com
 * @date 2022/11/26 16:49
 */

//输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
//
//
//
// 参考以下这颗二叉搜索树：
//
//      5
//    / \
//   2   6
//  / \
// 1   3
//
// 示例 1：
//
// 输入: [1,6,3,2,5]
//输出: false
//
// 示例 2：
//
// 输入: [1,3,2,6,5]
//输出: true
//
//
//
// 提示：
//
//
// 数组长度 <= 1000
//
//
// Related Topics 栈 树 二叉搜索树 递归 二叉树 单调栈 👍 637 👎 0

/*
	利用二叉搜索树的特性，最后一个为根节点，以根节点的值将原序列分为两部分，左边记为左子树，右边为右子树
	方法：
	1.根节点root的值为postorder[-1]
	2.遍历postorder，找到第一个大于root值的节点
		2.1 若无，则postorder[:len(postorder)-1]全为左子树,递归继续做即可；
		2.2 在index处找到第一个大于root值的节点，则postorder[:index]均小于root，为左子树，
			检验postorder[index:len(postorder)-1]是否都大于root，不是则return false；
			是则继续递归即可
*/
//leetcode submit region begin(Prohibit modification and deletion)
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	root := postorder[len(postorder)-1]
	index := len(postorder) - 1
	for i := 0; i < len(postorder); i++ {
		if postorder[i] > root {
			index = i
			break
		}
	}

	for i := index + 1; i < len(postorder); i++ {
		if postorder[i] < root {
			return false
		}
	}

	if index == len(postorder)-1 {
		return verifyPostorder(postorder[:index])
	}

	return verifyPostorder(postorder[:index]) && verifyPostorder(postorder[index:len(postorder)-1])
}

//leetcode submit region end(Prohibit modification and deletion)

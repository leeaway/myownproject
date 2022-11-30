package collections

/**
 * @author 2416144794@qq.com
 * @date 2022/11/25 15:13
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) IsLeaf() bool {
	return t != nil && t.Left == nil && t.Right == nil
}

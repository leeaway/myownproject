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

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

/*
	层序方式构建二叉树，-1表示null
	如：[1,2,3,4,5,-1,7]
*/
//func NewTreeNodeByArray(nums []int) *TreeNode {
//	if len(nums) == 0 || nums[0] == -1{
//		return nil
//	}
//	root := NewTreeNode(nums[0])
//	tmp := root
//	for i:=1;i<len(nums);i++ {
//
//	}
//}

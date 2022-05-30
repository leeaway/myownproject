package dailyExercise

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodeMap := make(map[*TreeNode]int)
	nodeMap[root] = root.Val
	que := []*TreeNode{root}
	res := 0
	for len(que) > 0 {
		curSize := len(que)
		for i := 0; i < curSize; i++ {
			top := que[0]
			que = que[1:]
			if top.Left != nil {
				nodeMap[top.Left] = nodeMap[top]*2 + top.Left.Val
				que = append(que, top.Left)
			}
			if top.Right != nil {
				nodeMap[top.Right] = nodeMap[top]*2 + top.Right.Val
				que = append(que, top.Right)
			}
			if top.Left == nil && top.Right == nil {
				res += nodeMap[top]
			}
		}
	}
	return res
}

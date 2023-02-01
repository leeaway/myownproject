package binarytree

import (
	"example.com/m/v2/tools/collections"
	"example.com/m/v2/tools/mathutil"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/1 11:21
 */

/*
示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 4, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
*/

/*
第一种只要求最长路径长度是多少，这个只要递归就可以：
	1. 因为最长路径一定等于某个节点的左右子树最大深度之和+1，所以我们递归每个节点，在中途记录最大的res即可
*/

var res int

func getMaxPathLen(root *collections.TreeNode) int {
	res = 0
	if root == nil {
		return 0
	}
	getMaxDepth(root)
	return res
}

func getMaxDepth(root *collections.TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := getMaxDepth(root.Left), getMaxDepth(root.Right)
	res = mathutil.MaxInt(res, l+r+1)
	return mathutil.MaxInt(l, r) + 1
}

/*
	若要找到具体的路径，则需要dfs求取，注意此时我们应该以叶子节点为起点开始遍历
	1. 第一步，因为子节点无法访问父节点，所以我们针对每个node，应该先维护好node所能到达的点neighbors（包括父节点，左右子节点）
	2. 开始dfs
*/

//层序遍历记录父节点
var longestPath []int

func buildNeighborsAndGetLeafs(root *collections.TreeNode) (map[*collections.TreeNode][]*collections.TreeNode, []*collections.TreeNode) {
	resMap := make(map[*collections.TreeNode][]*collections.TreeNode)
	longestPath = []int{}
	if root == nil {
		return resMap, nil
	}
	var leafNode []*collections.TreeNode
	q := []*collections.TreeNode{root}
	for len(q) > 0 {
		parent := q[0]
		if parent.IsLeaf() {
			leafNode = append(leafNode, parent)
		}
		q = q[1:]
		if parent.Left != nil {
			resMap[parent] = append(resMap[parent], parent.Left)
			resMap[parent.Left] = append(resMap[parent.Left], parent)
			q = append(q, parent.Left)
		}
		if parent.Right != nil {
			resMap[parent] = append(resMap[parent], parent.Right)
			resMap[parent.Right] = append(resMap[parent.Right], parent)
			q = append(q, parent.Right)
		}
	}
	return resMap, leafNode
}

func getLongestPath(root *collections.TreeNode) []int {
	if root == nil {
		return nil
	}
	neighbors, leafNode := buildNeighborsAndGetLeafs(root)
	for _, node := range leafNode {
		dfsHelper(node, neighbors, []int{}, make(map[*collections.TreeNode]bool))
	}
	return longestPath
}

func dfsHelper(start *collections.TreeNode, neighbors map[*collections.TreeNode][]*collections.TreeNode, path []int, visited map[*collections.TreeNode]bool) {
	if !canDown(start, neighbors, visited) {
		if !visited[start] {
			path = append(path, start.Val)
		}
		if len(path) > len(longestPath) {
			longestPath = path
		}
		return
	}
	if visited[start] {
		return
	}

	visited[start] = true
	path = append(path, start.Val)
	for _, neigh := range neighbors[start] {
		if visited[neigh] {
			continue
		}
		dfsHelper(neigh, neighbors, path, visited)
	}
	visited[start] = false
	path = path[:len(path)-1]
}

func canDown(node *collections.TreeNode, neighbors map[*collections.TreeNode][]*collections.TreeNode, visited map[*collections.TreeNode]bool) bool {
	for _, n := range neighbors[node] {
		if !visited[n] {
			return true
		}
	}
	return false
}

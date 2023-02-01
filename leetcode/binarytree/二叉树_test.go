package binarytree

import (
	"example.com/m/v2/tools/collections"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
 * @author 2416144794@qq.com
 * @date 2023/2/1 11:36
 */

func Test_getMaxPathLen(t *testing.T) {
	Convey("getMaxPathLen", t, func() {
		Convey("经过根节点", func() {
			root := collections.NewTreeNode(1)
			root.Left = collections.NewTreeNode(2)
			root.Right = collections.NewTreeNode(3)
			root.Left.Left = collections.NewTreeNode(4)
			root.Left.Right = collections.NewTreeNode(5)
			root.Right.Right = collections.NewTreeNode(7)
			So(getMaxPathLen(root), ShouldEqual, 5)
		})

		Convey("不经过根节点", func() {
			root := collections.NewTreeNode(1)
			root.Left = collections.NewTreeNode(2)
			root.Left.Left = collections.NewTreeNode(3)
			root.Left.Right = collections.NewTreeNode(4)
			root.Left.Left.Left = collections.NewTreeNode(5)
			root.Left.Right.Right = collections.NewTreeNode(6)
			root.Left.Right.Right.Right = collections.NewTreeNode(7)
			So(getMaxPathLen(root), ShouldEqual, 6)
		})

		Convey("jiandan", func() {
			root := collections.NewTreeNode(1)
			root.Left = collections.NewTreeNode(2)
			So(getMaxPathLen(root), ShouldEqual, 2)
		})
	})
}

func Test_buildNeighbors(t *testing.T) {
	Convey("buildNeighbors", t, func() {
		root := collections.NewTreeNode(1)
		root.Left = collections.NewTreeNode(2)
		root.Right = collections.NewTreeNode(3)
		root.Left.Left = collections.NewTreeNode(4)
		root.Left.Right = collections.NewTreeNode(5)
		root.Right.Right = collections.NewTreeNode(7)
		neighbors, leafNode := buildNeighborsAndGetLeafs(root)
		for _, node := range leafNode {
			fmt.Printf("%d ", node.Val)
		}
		fmt.Println()
		for k, v := range neighbors {
			fmt.Printf("%d:", k.Val)
			for _, node := range v {
				fmt.Printf("%d,", node.Val)
			}
			fmt.Println()
		}
	})
}

func Test_dfsHelper(t *testing.T) {
	Convey("dfsHelper", t, func() {
		Convey("dfsHelper test1", func() {
			root := collections.NewTreeNode(1)
			root.Left = collections.NewTreeNode(2)
			root.Right = collections.NewTreeNode(3)
			root.Left.Left = collections.NewTreeNode(4)
			root.Left.Right = collections.NewTreeNode(5)
			root.Right.Right = collections.NewTreeNode(7)
			fmt.Println(getLongestPath(root))
		})
	})
}

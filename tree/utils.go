package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// GenerateTreeFromLevelOrder 从层序遍历构造二叉树
func GenerateTreeFromLevelOrder(orders []int) *TreeNode {
	treeNodes := make([]*TreeNode, len(orders))
	for i := 0; i < len(orders); i++ {
		if orders[i] != -1 {
			treeNodes[i] = &TreeNode{Val: orders[i]}
		}
	}
	for i := 0; 2*i+1 < len(orders); i++ {
		if treeNodes[i] != nil {
			treeNodes[i].Left = treeNodes[2*i+1]
			if 2*i+2 < len(orders) {
				treeNodes[i].Right = treeNodes[2*i+2]
			}
		}
	}
	return treeNodes[0]
}

func PrintAllPathsToLeaves(root *TreeNode) {
	ans := make([]string, 0)
	var dfs func(node *TreeNode, path []string)
	dfs = func(node *TreeNode, path []string) {
		if node == nil {
			return
		}
		path = append(path, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil { //叶子结点
			ans = append(ans, strings.Join(path, "=>"))
			return
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, []string{})
	fmt.Println(ans)
}

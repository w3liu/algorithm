package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	ok := isValidBST(root)
	fmt.Println("ok", ok)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return compare(root, math.MinInt64, math.MaxInt64)
}

func compare(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return compare(root.Left, lower, root.Val) && compare(root.Right, root.Val, upper)
}

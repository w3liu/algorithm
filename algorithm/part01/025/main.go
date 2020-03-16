package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	root := pruneTree(node)
	b, e := json.Marshal(root)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

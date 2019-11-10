package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 11,
				},
				Right: &TreeNode{
					Val: 12,
				},
			},
			Right: &TreeNode{
				Val: 13,
				Left: &TreeNode{
					Val: 14,
				},
				Right: &TreeNode{
					Val: 15,
				},
			},
		},
	}
	results := inorderTraversal(root)
	fmt.Println(results)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	results := make([]int, 0)
	if root.Left != nil {
		results = append(results, inorderTraversal(root.Left)...)
	}

	results = append(results, root.Val)

	if root.Right != nil {
		results = append(results, inorderTraversal(root.Right)...)
	}
	return results
}

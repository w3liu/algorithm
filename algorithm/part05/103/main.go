package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root = &TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			preLeft := preorder[1 : i+1]
			preRight := preorder[i+1:]
			inLeft := inorder[0:i]
			inRight := inorder[i+1:]
			root.Left = buildTree(preLeft, inLeft)
			root.Right = buildTree(preRight, inRight)
			break
		}
	}
	return root
}

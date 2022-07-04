package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ret = 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		ret = max(ret, l+r+1)
		return max(l, r) + 1
	}
	depth(root)
	return ret - 1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxNum int

func maxPathSum(root *TreeNode) int {
	maxNum = math.MinInt64
	maxNodeSum(root)
	return maxNum
}

func maxNodeSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftNum := int(math.Max(float64(maxNodeSum(node.Left)), 0))
	rightNum := int(math.Max(float64(maxNodeSum(node.Right)), 0))

	newNum := node.Val + leftNum + rightNum

	if maxNum < newNum {
		maxNum = newNum
	}
	return node.Val + int(math.Max(float64(leftNum), float64(rightNum)))
}

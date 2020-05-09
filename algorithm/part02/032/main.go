package main

import "fmt"

/**
[3,4,5,1,2,null,null,0]
[4,1,2]
[3,4,5,1,null,2]
[3,1,2]
*/

func main() {
	//s := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 1,
	//		},
	//		Right: &TreeNode{
	//			Val: 2,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 5,
	//	},
	//}
	//t := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//	},
	//}
	s := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
		},
		//Right: &TreeNode{
		//	Val: 1,
		//},
	}
	t := &TreeNode{
		Val: 1,
	}
	b := isSubtree(s, t)
	fmt.Println("b", b)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return check(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func check(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if (s != nil && t == nil) || (s == nil && t != nil) || (s.Val != t.Val) {
		return false
	}
	return check(s.Left, t.Left) && check(s.Right, t.Right)
}

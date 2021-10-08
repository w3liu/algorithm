package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var height = 1
	var result = make([][]int, 0)
	return render(root, height, result)
}

func render(node *TreeNode, height int, result [][]int) [][]int {
	if node == nil {
		return result
	}
	if len(result) < height {
		result = append(result, []int{})
	}
	result[height-1] = append(result[height-1], node.Val)
	result = render(node.Left, height+1, result)
	result = render(node.Right, height+1, result)
	return result
}

func bfs(root *TreeNode) [][]int {
	var result = make([][]int, 0)
	if root == nil {
		return result
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var tmp []*TreeNode
		var ret []int
		for i := 0; i < len(queue); i++ {
			ret = append(ret, queue[i].Val)
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		result = append(result, ret)
		queue = tmp
	}
	return result
}

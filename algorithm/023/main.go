package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	post := []int{4, 5, 2, 6, 7, 3, 1}
	node := constructFromPrePost(pre, post)
	data, _ := json.Marshal(node)
	fmt.Println(string(data))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
我们令左分支有 L 个节点。我们知道左分支的头节点为 pre[1]，但它也出现在左分支的后序表示的最后。
所以 pre[1] = post[L-1]（因为结点的值具有唯一性），因此 L = post.indexOf(pre[1]) + 1。

现在在我们的递归步骤中，左分支由 pre[1 : L+1] 和 post[0 : L] 重新分支，而右分支将由 pre[L+1 : N] 和 post[L : N-1] 重新分支。
*/

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: pre[0],
	}
	if len(pre) == 1 {
		return root
	}
	l := 0
	for i, v := range post {
		if v == pre[1] {
			l = i + 1
			break
		}
	}
	left := constructFromPrePost(pre[1:l+1], post[0:l])
	right := constructFromPrePost(pre[l+1:], post[l:len(post)-1])
	root.Left = left
	root.Right = right
	return root
}

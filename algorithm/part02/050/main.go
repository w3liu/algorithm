package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	iterator := Constructor(root)
	fmt.Println(iterator.Next())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
	fmt.Println(iterator.Next())
	fmt.Println(iterator.HasNext())
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	queue *list.List
}

func Constructor(root *TreeNode) BSTIterator {
	queue := list.New()
	midOrder(root, queue)
	return BSTIterator{
		queue: queue,
	}
}

func midOrder(node *TreeNode, queue *list.List) {
	for node != nil {
		queue.PushFront(node)
		node = node.Left
	}
}

func (bst *BSTIterator) Next() int {
	elem := bst.queue.Front()
	bst.queue.Remove(elem)
	node := elem.Value.(*TreeNode)
	if node.Right != nil {
		midOrder(node.Right, bst.queue)
	}
	return node.Val
}

func (bst *BSTIterator) HasNext() bool {
	return bst.queue.Len() > 0
}

/**
Your BSTIterator object will be instantiated and called as such:
obj := Constructor(root);
param_1 := obj.Next();
param_2 := obj.HasNext();
*/

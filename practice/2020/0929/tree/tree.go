package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &Node{
		Value: 'A',
		Left: &Node{
			Value: 'B',
			Left: &Node{
				Value: 'D',
				Left: &Node{
					Value: 'H',
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 'I',
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: 'E',
				Left: &Node{
					Value: 'J',
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		Right: &Node{
			Value: 'C',
			Left: &Node{
				Value: 'F',
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Value: 'G',
				Left:  nil,
				Right: nil,
			},
		},
	}
	arr := levelRange(root)
	for _, item := range arr {
		fmt.Println(string(item))
	}
}

type Node struct {
	Value byte
	Left  *Node
	Right *Node
}

// 二叉树前序遍历
func preOrder(root *Node) []byte {
	if root == nil {
		return []byte{}
	}
	arr := make([]byte, 0)
	arr = append(arr, root.Value)
	arr = append(arr, preOrder(root.Left)...)
	arr = append(arr, preOrder(root.Right)...)
	return arr
}

// 二叉树中序遍历
func inOrder(root *Node) []byte {
	if root == nil {
		return []byte{}
	}
	arr := make([]byte, 0)
	arr = append(arr, inOrder(root.Left)...)
	arr = append(arr, root.Value)
	arr = append(arr, inOrder(root.Right)...)
	return arr
}

// 二叉树后续遍历
func postOrder(root *Node) []byte {
	if root == nil {
		return []byte{}
	}
	arr := make([]byte, 0)
	arr = append(arr, postOrder(root.Left)...)
	arr = append(arr, postOrder(root.Right)...)
	arr = append(arr, root.Value)
	return arr
}

// 二叉树按层遍历
func levelRange(root *Node) []byte {
	if root == nil {
		return []byte{}
	}
	arr := make([]byte, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		front := queue.Front()
		queue.Remove(front)
		n := front.Value.(*Node)
		arr = append(arr, n.Value)
		if n.Left != nil {
			queue.PushBack(n.Left)
		}
		if n.Right != nil {
			queue.PushBack(n.Right)
		}
	}
	return arr
}

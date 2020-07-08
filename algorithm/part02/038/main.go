package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	node = reverseList(node)
	fmt.Println(node.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	var prev *ListNode
	for head != nil {
		curr = head
		head = head.Next
		curr.Next = prev
		prev = curr
	}
	return curr
}

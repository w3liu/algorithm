package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	node := sortList(head)
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	h := &ListNode{
		Val:  0,
		Next: nil,
	}
	res := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}
	return res.Next
}

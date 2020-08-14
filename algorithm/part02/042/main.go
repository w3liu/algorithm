package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := middle(head)
	left := head
	right := mid.Next
	mid.Next = nil

	right = reverse(right)

	merge(left, right)
}

func middle(head *ListNode) *ListNode {
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		temp := &ListNode{
			Val:  curr.Val,
			Next: prev,
		}
		prev = temp
		curr = curr.Next
	}
	return prev
}

func merge(left, right *ListNode) {
	for left != nil && right != nil {
		leftTemp := left.Next
		rightTemp := right.Next

		left.Next = right
		left = leftTemp

		right.Next = leftTemp
		right = rightTemp
	}
}

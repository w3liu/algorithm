package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	b := insertionSortList(node)
	for b != nil {
		fmt.Println(b.Val)
		b = b.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	curNode := head.Next
	headNode := head
	tailNode := head

	for curNode != nil {
		tempNode := curNode.Next

		if headNode.Val > curNode.Val {
			curNode.Next = headNode
			headNode = curNode
		} else if tailNode.Val <= curNode.Val {
			tailNode.Next = curNode
			tailNode = curNode
		} else {
			seqNode := headNode
			for seqNode != tailNode {
				if seqNode.Next.Val > curNode.Val {
					curNode.Next = seqNode.Next
					seqNode.Next = curNode
					break
				} else {
					seqNode = seqNode.Next
				}
			}
		}
		curNode = tempNode
	}
	tailNode.Next = nil
	return headNode
}

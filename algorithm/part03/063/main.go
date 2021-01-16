package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var l3 = new(ListNode)
	var head = &ListNode{
		Next: l3,
	}
	var v1, v2, v3, next int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		v3 = v1 + v2 + next
		if v3 >= 10 {
			next = v3 / 10
			v3 = v3 % 10
		} else {
			next = 0
		}
		l3.Next = &ListNode{
			Val: v3,
		}
		l3 = l3.Next
	}
	if next > 0 {
		l3.Next = &ListNode{
			Val: next,
		}
	}
	return head.Next.Next
}

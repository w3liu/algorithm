package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	l3 := addTwoNumbers(l1, l2)
	num := ""
	for l3 != nil {
		val := l3.Val
		num = fmt.Sprintf("%s%d", num, val)
		l3 = l3.Next
	}
	fmt.Println(num)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var previous *ListNode
	var current *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		result := v1 + v2 + carry
		nodeVal := result % 10
		carry = result / 10
		current = &ListNode{
			Val: nodeVal,
		}
		if head == nil {
			head = current
		}
		if previous != nil {
			previous.Next = current
		}

		previous = current

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry > 0 {
		current = &ListNode{
			Val: carry,
		}
		if head == nil {
			head = current
		}
		if previous != nil {
			previous.Next = current
		}
	}

	return head
}

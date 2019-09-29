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
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l3 := mergeTwoLists(l1, l2)
	fmt.Println("l3", l3)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	var pre *ListNode
	var curr *ListNode
	var i int
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val >= l2.Val {
				curr = &ListNode{
					Val: l2.Val,
				}
				l2 = l2.Next
			} else {
				curr = &ListNode{
					Val: l1.Val,
				}
				l1 = l1.Next
			}
		} else if l1 == nil {
			curr = l2
			l2 = l2.Next
		} else {
			curr = l1
			l1 = l1.Next
		}
		if i == 0 {
			l3 = curr
		} else {
			pre.Next = curr
		}
		pre = curr
		i++
	}
	return l3
}

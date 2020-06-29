package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					//Next: &ListNode{
					//	Val:  1,
					//	Next: nil,
					//},
				},
			},
		},
	}
	result := isPalindrome(list)
	fmt.Println("result", result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 慢指针
	p1 := head
	// 快指针
	p2 := head
	var prev, left, right *ListNode
	for true {
		if p2 == nil {
			break
		}
		if p2.Next == nil {
			left = p1
			right = &ListNode{
				Val:  p1.Val,
				Next: p1.Next,
			}
			left.Next = prev
			break
		} else {
			// 快指针每次移动两个节点
			p2 = p2.Next.Next
			left = p1
			p1 = p1.Next
			right = p1
			// 修改左边链表的next,使其反向链接
			left.Next = prev
			prev = left
		}
	}
	// 对比左右链表
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

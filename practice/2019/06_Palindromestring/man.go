package main

import "fmt"

func main() {
	node := &Node{
		Val: 'A',
		Next: &Node{
			Val: 'B',
			Next: &Node{
				Val: 'C',
				Next: &Node{
					Val: 'D',
					Next: &Node{
						Val: 'C',
						Next: &Node{
							Val: 'B',
							Next: &Node{
								Val:  'A',
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	b := judgePalindrome(node)
	fmt.Println(b)
}

type Node struct {
	Val  byte
	Next *Node
}

func judgePalindrome(node *Node) bool {
	var prev *Node
	fast := node
	slow := node
	// 找中点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	if fast != nil {
		slow = slow.Next
	}
	// 比较
	for slow != nil && prev != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow = slow.Next
		prev = prev.Next
	}
	return true
}

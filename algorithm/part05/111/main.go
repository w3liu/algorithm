package main

import "fmt"

func main() {
	var nodes = []int{1, 2, 3, 4}
	var list = createList(nodes)
	var curr = list
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	// 制造环
	list.Next.Next.Next = list
	var has = hasCycle(list)
	fmt.Println(has)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil ||
		head.Next.Next == nil {
		return false
	}
	var p1 = head
	var p2 = head.Next.Next
	for true {
		if p1 == nil || p2 == nil || p2.Next == nil {
			break
		}
		if p1 == p2 {
			return true
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return false
}

func createList(nodes []int) *ListNode {
	var head = &ListNode{}
	var curr = head
	for _, v := range nodes {
		tmp := &ListNode{Val: v}
		curr.Next = tmp
		curr = tmp
	}
	return head.Next
}

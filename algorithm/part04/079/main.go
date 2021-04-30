package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p = &ListNode{Next: head}
	var l = head
	var r = head
	var cnt int
	for r != nil {
		r = r.Next
		if cnt < n+1 {
			cnt++
		} else {
			l = l.Next
		}
	}
	if cnt < n+1 {
		p.Next = p.Next.Next
		return p.Next
	} else {
		l.Next = l.Next.Next
		return head
	}
}

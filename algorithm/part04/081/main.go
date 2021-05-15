package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	for _, node := range lists {
		result = merge(result, node)
	}
	return result
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head = &ListNode{}
	var tail, p1, p2 = head, l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tail.Next = p1
			p1 = p1.Next
		} else {
			tail.Next = p2
			p2 = p2.Next
		}
		tail = tail.Next
	}
	if p1 != nil {
		tail.Next = p1
	} else {
		tail.Next = p2
	}
	return head.Next
}

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	isCircle := false
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			isCircle = true
			break
		}
	}
	if isCircle {
		ptr1 := head
		ptr2 := slow
		for {
			if ptr1 == ptr2 {
				return ptr1
			}
			ptr1 = ptr1.Next
			ptr2 = ptr2.Next
		}
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]int)
	node := head
	for {
		if node == nil {
			return nil
		}
		if _, ok := nodeMap[node]; ok {
			return node
		}
		nodeMap[node] = node.Val
		node = node.Next
	}
}

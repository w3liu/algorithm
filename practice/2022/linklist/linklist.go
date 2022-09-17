package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var curr, prev *ListNode
	for head != nil {
		curr = head
		head = head.Next
		curr.Next = prev
		prev = curr
	}
	return curr
}

func checkCircle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	var slow, quik *ListNode
	slow = head
	quik = head.Next.Next
	for quik != nil {
		if slow == quik {
			return true
		}
		if quik.Next == nil || quik.Next.Next == nil {
			return false
		}
		quik = quik.Next.Next
		slow = slow.Next
	}
	return false
}

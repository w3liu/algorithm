package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
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

// 检查环
func checkCircle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	var slow, quick *ListNode
	slow = head
	quick = head.Next.Next
	for quick != nil {
		if slow == quick {
			return true
		}
		if quick.Next == nil || quick.Next.Next == nil {
			return false
		}
		quick = quick.Next.Next
		slow = slow.Next
	}
	return false
}

// 合并链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, tail, curr *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curr = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		} else {
			curr = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		}
		if tail == nil {
			head = curr
		} else {
			tail.Next = curr
		}
		tail = curr
	}
	if list1 != nil {
		if tail != nil {
			tail.Next = list1
		} else {
			return list1
		}
	}
	if list2 != nil {
		if tail != nil {
			tail.Next = list2
		} else {
			return list2
		}
	}
	return head
}

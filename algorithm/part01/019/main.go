package main

import "fmt"

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	b := insertionSortList(node)
	for b != nil {
		fmt.Println(b.Val)
		b = b.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	origin := &ListNode{Val: 0, Next: head}
	curr := head
	tail := head
	for curr != nil {
		next := curr.Next
		// 当前节点和尾部节点相遇，不做任何处理，直接进入下一轮
		if tail == curr {
			curr = next
			continue
		}
		// 若干当前节点值大于尾部节点，尾部节点与当前节点做交换
		if curr.Val >= tail.Val {
			tail.Next = curr
			tail = curr
		} else {
			// 取出head节点的第一个节点
			seq := origin.Next
			// 取出origin节点的第一个节点（head的前置节点）
			pre := origin
			// 从头到尾开始遍历节点
			for seq != tail {
				// 如果遇到有大于当前节点的，则中断
				if seq.Val > curr.Val {
					break
				}
				pre = seq
				seq = seq.Next
			}
			// 执行插入排序
			curr.Next = seq
			pre.Next = curr
		}

		curr = next
	}
	// 尾部的Next需要置空，要不可能会出现死循环
	tail.Next = nil
	return origin.Next
}

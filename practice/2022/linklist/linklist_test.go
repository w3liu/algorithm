package linklist

import "testing"

func TestReverseList(t *testing.T) {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	ret := reverseList(list)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

func TestCheckCircle(t *testing.T) {
	node4 := &ListNode{}
	node4.Next = &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  6,
				Next: node4,
			},
		},
	}
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: node4,
			},
		},
	}
	hasCircle := checkCircle(list)
	t.Log(hasCircle)
}

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	ret := mergeTwoLists(list1, list2)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	}
	ret := removeNthFromEnd(list1, 2)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

func TestRemoveNthFromEnd1(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	ret := removeNthFromEnd(list1, 2)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

func TestMiddleNode(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	ret := middleNode(list1)
	for ret != nil {
		t.Log(ret.Val)
		ret = ret.Next
	}
}

func TestDetectCycle(t *testing.T) {
	node4 := &ListNode{
		Val: 4,
	}
	node4.Next = &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  7,
				Next: node4,
			},
		},
	}
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: node4,
			},
		},
	}
	ret := detectCycle(list)
	if ret != nil {
		t.Log(ret.Val)
	}
}

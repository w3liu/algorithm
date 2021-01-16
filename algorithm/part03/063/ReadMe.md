## Algorithm
### 1. 题目
```
2. 两数相加
```
### 2. 题目描述
```
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
```

### 3. 解答：
```golang
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var l3 = new(ListNode)
	var head = &ListNode{
		Next: l3,
	}
	var v1, v2, v3, next int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		v3 = v1 + v2 + next
		if v3 >= 10 {
			next = v3 / 10
			v3 = v3 % 10
		} else {
			next = 0
		}
		l3.Next = &ListNode{
			Val: v3,
		}
		l3 = l3.Next
	}
	if next > 0 {
		l3.Next = &ListNode{
			Val: next,
		}
	}
	return head.Next.Next
}
```
### 4. 说明
1. 每个节点相加如果有进位，就保存到next变量
2. 如果最后next变量不为0则还需要增加一个节点

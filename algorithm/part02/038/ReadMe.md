## Algorithm
### 1. 题目
```
206. 反转链表
```
### 2. 题目描述
```
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
```

### 3. 解答：
```golang
func reverseList(head *ListNode) *ListNode {
	var curr *ListNode
	var prev *ListNode
	for head != nil {
		curr = head
		head = head.Next
		curr.Next = prev
		prev = curr
	}
	return curr
}
```
### 4. 说明
双指针，curr指针反转，prev记录上一个节点
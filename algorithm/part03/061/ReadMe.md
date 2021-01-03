## Algorithm
### 1. 题目
```
24. 两两交换链表中的节点
```
### 2. 题目描述
```
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 

示例 1：

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]

```

### 3. 解答：
```golang
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	head.Next = swapPairs(result.Next)
	result.Next = head
	return result
}
```
### 4. 说明
采用递归，终止条件是head为空或者head.Next为空
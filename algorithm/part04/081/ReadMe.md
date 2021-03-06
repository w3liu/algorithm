## Algorithm
### 1. 题目
```
23. 合并K个升序链表
```
### 2. 题目描述
```
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

 

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]
 

提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

```

### 3. 解答：
```golang
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
```
### 4. 说明
1. 定义一个合并两个链表的方法；
2. 遍历所有链表进行合并；
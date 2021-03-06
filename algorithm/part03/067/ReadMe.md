## Algorithm
### 1. 题目
```
160. 相交链表
```
### 2. 题目描述
```
编写一个程序，找到两个单链表相交的起始节点。

如下面的两个链表：
```

[!qr](./images/0214_a_1.png)

```
在节点 c1 开始相交。
```

### 3. 解答：
```golang
type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		if pA == pB {
			return pA
		}
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
```
### 4. 说明
1. 采用双指针pA pB 分别指向 headA headB
2. 循环移动指针pA与pB，当pA不等于pB的时候终止
3. 终止的情况有两种，一种是遇到相交节点，pA == pB，另外一种是遍历了m+n此没有相交节点，大家都是nil
4. 这道题的核心思想是先后遍历两个链表，最终遍历的长度是一致的，都是m+n
## Algorithm
### 1. 题目
```
19. 删除链表的倒数第 N 个结点
```
### 2. 题目描述
```
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
 

提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

```

### 3. 解答：
```golang
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p = &ListNode{Next: head}
	var l = head
	var r = head
	var cnt int
	for r != nil {
		r = r.Next
		if cnt < n+1 {
			cnt++
		} else {
			l = l.Next
		}
	}
	if cnt < n+1 {
		p.Next = p.Next.Next
		return p.Next
	} else {
		l.Next = l.Next.Next
		return head
	}
}
```
### 4. 说明
1. 定义一个节点p，next指针指向head；
2. 定义l,r指向head;
3. 定义cnt记录l与r之间的距离；
4. 首先将r指针向右移动，同时增加cnt；
5. 当cnt与n+1相等的时候，同时移动l，r直到r==nil；
6. 最后如果cnt < n+1则返回p.next，除此之外，返回head；
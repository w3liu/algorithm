## Algorithm
### 1. 题目
```
199. 二叉树的右视图
```
### 2. 题目描述
```
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---

```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := list.New()
	result := make([]int, 0)
	queue.PushFront(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			elem := queue.Front()
			queue.Remove(elem)
			node := elem.Value.(*TreeNode)
			if i == 0 {
				result = append(result, node.Val)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
		}

	}
	return result
}
```
### 4. 说明

1. 定义一个队列用于存放每一层的节点
2. queque.len()表示一层总共有多少节点
3. 遍历这一层的全部节点，将最右边的节点存入结果数组
4. 接着先将有节点存入队列，再将左节点存入结果数组
5. 最后返回结果数组
## Algorithm
### 1. 题目
```
114. 二叉树展开为链表
```
### 2. 题目描述
```
给定一个二叉树，原地将它展开为一个单链表。

 

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			prev := next
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
```
### 4. 说明
1. 对于当前节点，如果其左子节点不为空，则在其左子树中找到最右边的节点，作为前驱节点。
2. 将当前节点的右子节点赋给前驱节点的右子节点，然后将当前节点的左子节点赋给当前节点的右子节点，并将当前节点的左子节点设为空。
3. 对当前节点处理结束后，继续处理链表中的下一个节点，直到所有节点都处理结束。
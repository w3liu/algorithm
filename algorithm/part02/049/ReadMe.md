## Algorithm
### 1. 题目
```
104. 二叉树的最大深度
```
### 2. 题目描述
```
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := maxDepth(root.Right) + 1
	left := maxDepth(root.Left) + 1
	if right > left {
		return right
	}
	return left
}
```
### 4. 说明
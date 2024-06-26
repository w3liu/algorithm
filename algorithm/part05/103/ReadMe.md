## Algorithm
### 1. 题目
```
105. 从前序与中序遍历序列构造二叉树
```
### 2. 题目描述
```
给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。

```

[!qr](./images/1011_a_1.jpg)

```
示例 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
示例 2:

Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var root = &TreeNode{Val: preorder[0]}
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			preLeft := preorder[1 : i+1]
			preRight := preorder[i+1:]
			inLeft := inorder[0:i]
			inRight := inorder[i+1:]
			root.Left = buildTree(preLeft, inLeft)
			root.Right = buildTree(preRight, inRight)
			break
		}
	}
	return root
}
```
### 4. 说明
采用递归的方式，当preorder或者inorder为空的时候就返回。
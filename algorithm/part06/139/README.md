## Algorithm
### 1. 题目
```
543. 二叉树的直径
```
### 2. 题目描述
```
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

 

示例 :
给定二叉树

          1
         / \
        2   3
       / \     
      4   5    
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ret = 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		ret = max(ret, l+r+1)
		return max(l, r) + 1
	}
	depth(root)
	return ret - 1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
```
### 4. 说明
采用递归深度优先遍历算法，记录每个节点的最大深度以及该节点为根的节点总和。
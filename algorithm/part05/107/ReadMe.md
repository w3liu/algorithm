## Algorithm
### 1. 题目
```
124. 二叉树中的最大路径和
```
### 2. 题目描述
```
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
```

示例1：

[!qr](./images/1108_a_1.jpg)

```
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
```

示例2：

[!qr](./images/1108_a_2.jpg)

```
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxNum int

func maxPathSum(root *TreeNode) int {
	maxNum = math.MinInt64
	maxNodeSum(root)
	return maxNum
}

func maxNodeSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftNum := int(math.Max(float64(maxNodeSum(node.Left)), 0))
	rightNum := int(math.Max(float64(maxNodeSum(node.Right)), 0))

	newNum := node.Val + leftNum + rightNum

	if maxNum < newNum {
		maxNum = newNum
	}
	return node.Val + int(math.Max(float64(leftNum), float64(rightNum)))
}
```
### 4. 说明
递归
## Algorithm
### 1. 题目
```
102. 二叉树的层序遍历
```
### 2. 题目描述
```
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

```

### 3. 解答：
```golang
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var height = 1
	var result = make([][]int, 0)
	return render(root, height, result)
}

func render(node *TreeNode, height int, result [][]int) [][]int {
	if node == nil {
		return result
	}
	if len(result) < height {
		result = append(result, []int{})
	}
	result[height-1] = append(result[height-1], node.Val)
	result = render(node.Left, height+1, result)
	result = render(node.Right, height+1, result)
	return result
}

func bfs(root *TreeNode) [][]int {
	var result = make([][]int, 0)
	if root == nil {
		return result
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var tmp []*TreeNode
		var ret []int
		for i := 0; i < len(queue); i++ {
			ret = append(ret, queue[i].Val)
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		result = append(result, ret)
		queue = tmp
	}
	return result
}
```
### 4. 说明
采用了两种方式：
    * 第一种是深度优先遍历，递归给result数组赋值；
    * 第二种是广度优先遍历，采用先进先出队列按层给result数组赋值；

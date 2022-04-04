## Algorithm
### 1. 题目
```
337. 打家劫舍 III
```
### 2. 题目描述
```
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
```

示例1：

[!qr](./images/0404_a_1.jpg)

```
输入: root = [3,2,3,null,3,null,1]
输出: 7 
解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
```

示例2：

[!qr](./images/0404_a_2.jpg)

```
输入: root = [3,4,5,1,3,null,1]
输出: 9
解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
```

### 3. 解答：
```golang
func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	// if select current node, then left and right node can not be selected
	selected := node.Val + l[1] + r[1]
	// if current node not select, left and right can be selected or not
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```
### 4. 说明
使用动态规划，状态转移方程：
当前节点选中：f(n) = g(l) + g(r)
当前节点未选中： g(n)=max{f(l),g(l)}+max{f(r),g(r)}
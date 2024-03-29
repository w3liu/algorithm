## Algorithm
### 1. 题目
```
62. 不同路径
```
### 2. 题目描述
```
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

示例1
```
[!qr](./images/0726_a_1.png)

```
输入：m = 3, n = 7
输出：28
```
```
示例 2：

输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下

示例 3：

输入：m = 7, n = 3
输出：28
示例 4：

输入：m = 3, n = 3
输出：6
```
### 3. 解答：
```golang
func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			dp[x][y] = dp[x-1][y] + dp[x][y-1]
		}
	}
	return dp[m-1][n-1]
}
```
### 4. 说明
动态规划，动态转移方程：`f(i,j) = f(i-1,j) + f(i, j-1)`
## Algorithm
### 1. 题目
```
64. 最小路径和
```
### 2. 题目描述
```
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

 

示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12

```

### 3. 解答：
```golang
func minPathSum(grid [][]int) int {
	var m, n = len(grid), len(grid[0])
	var dp = make([][]int, len(grid))
	dp[0] = make([]int, len(grid[0]))
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, len(grid[0]))
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			dp[x][y] = int(math.Min(float64(dp[x-1][y]), float64(dp[x][y-1]))) + grid[x][y]
		}
	}
	return dp[m-1][n-1]
}
```
### 4. 说明
动态规划，dp用来存放最小路径的值，第一行和第一列需要特殊处理，最终返回`dp[m-1][n-1]`即可
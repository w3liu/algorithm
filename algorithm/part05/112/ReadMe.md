## Algorithm
### 1. 题目
```
221. 最大正方形
```
### 2. 题目描述
```
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
```
示例 1：
[!qr](./images/1213_a_1.jpg)
```
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4
```
示例 2：
[!qr](./images/1213_a_2.jpg)
```
输入：matrix = [["0","1"],["1","0"]]
输出：1
```
示例 3：
```
输入：matrix = [["0"]]
输出：0
```

### 3. 解答：
```golang
func maximalSquare(matrix [][]byte) int {
	var maxSide int
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return maxSide
	}
	var rows = len(matrix)
	var cols = len(matrix[0])
	var dp = make([][]int, 0, rows)
	for i := 0; i < rows; i++ {
		dp = append(dp, make([]int, cols))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]),
						float64(dp[i-1][j-1])),
						float64(dp[i][j-1]))) + 1
				}
			} else {
				dp[i][j] = 0
			}
			maxSide = int(math.Max(float64(maxSide), float64(dp[i][j])))
		}
	}
	return maxSide * maxSide
}
```
### 4. 说明
采用动态规划，设置最长变成为maxSide，`dp[i][j]`表示以(i,j)为右下角，且只包含1的正方形的边长最大值。
状态转移方程如下：
dp(i,j) = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
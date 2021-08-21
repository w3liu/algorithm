## Algorithm
### 1. 题目
```
72. 编辑距离
```
### 2. 题目描述
```
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
 

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
```

### 3. 解答：
```golang
func minDistance(word1 string, word2 string) int {
	var m, n = len(word1), len(word2)
	if m == 0 || n == 0 {
		return m + n
	}
	var dp = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + int(math.Min(float64(dp[i-1][j]), math.Min(float64(dp[i][j-1]), float64(dp[i-1][j-1]))))
			}
		}
	}
	return dp[m][n]
}
```
### 4. 说明
* 如果`word1[i] == word2[j]`, 那么`dp[i][j] = dp[i-1][j-1]`
* 否则，`dp[i][j] = 1 + min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1])`